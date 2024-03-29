package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
	distrKeeper   types.DistrKeeper

	paramstore paramtypes.Subspace
}

// NewKeeper creates a new allocation Keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, stakingKeeper types.StakingKeeper, distrKeeper types.DistrKeeper,
	ps paramtypes.Subspace,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		accountKeeper: accountKeeper, bankKeeper: bankKeeper, stakingKeeper: stakingKeeper,
		distrKeeper: distrKeeper,
		paramstore:  ps,
	}
}

// DistributeInflation distributes module-specific inflation
func (k Keeper) DistributeInflation(ctx sdk.Context) error {
	// Fee collector module account account contains newly minted coins and collected fees from transactions
	blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).GetAddress()
	blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, k.stakingKeeper.BondDenom(ctx))

	// if there is no inflation or fees, return
	if blockInflation.IsZero() {
		return nil
	}
	params := k.GetParams(ctx)
	proportions := params.DistributionProportions

	// fund randomness rewards address
	if params.RandomnessRewardsReceiver != "" {
		randomnessRewardsCoin := k.GetProportions(ctx, blockInflation, proportions.RandomnessRewards)
		randomnessRewardsReceiver, err := sdk.AccAddressFromBech32(params.RandomnessRewardsReceiver)
		if err != nil {
			return err
		}
		if !randomnessRewardsCoin.IsZero() {
			err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, authtypes.FeeCollectorName, randomnessRewardsReceiver, sdk.NewCoins(randomnessRewardsCoin))
			if err != nil {
				return err
			}
		}
	}

	// fund validator rewards pool
	validatorRewardsCoins := k.GetProportions(ctx, blockInflation, proportions.ValidatorRewards)
	err := k.DistributeValidatorRewards(ctx, validatorRewardsCoins)
	if err != nil {
		return err
	}
	devRewards := k.GetProportions(ctx, blockInflation, proportions.DeveloperRewards)
	err = k.DistributeDeveloperRewards(ctx, blockInflationAddr, devRewards, params.WeightedDeveloperRewardsReceivers)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) DistributeDeveloperRewards(ctx sdk.Context, feeCollectorAddress sdk.AccAddress, devRewards sdk.Coin, devs []types.WeightedAddress) error {
	if devRewards.IsZero() {
		return nil
	}
	for _, w := range devs {
		devRewardPortionCoins := sdk.NewCoins(k.GetProportions(ctx, devRewards, w.Weight))
		if w.Address != "" {
			devRewardsAddr, err := sdk.AccAddressFromBech32(w.Address)
			if err != nil {
				return err
			}
			err = k.bankKeeper.SendCoins(ctx, feeCollectorAddress, devRewardsAddr, devRewardPortionCoins)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DistributeValidatorRewards distributes rewards to validators
func (k Keeper) DistributeValidatorRewards(ctx sdk.Context, rewards sdk.Coin) error {
	if rewards.IsZero() {
		return nil
	}
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, authtypes.FeeCollectorName, types.ValidatorRewardsPool, sdk.NewCoins(rewards))
	if err != nil {
		return err
	}
	validators := k.stakingKeeper.GetLastValidators(ctx)
	// this should never happen but adding the validation
	if len(validators) == 0 {
		return nil
	}
	// get the amount of coins to distribute to each validator
	validatorReward := rewards.Amount.QuoRaw(int64(len(validators)))
	// distribute coins to each validator by accumulating their rewards
	// the module account will hold the tokens until they are withdrawn by validators
	for _, v := range validators {
		// get the validator operator address
		operator, err := sdk.ValAddressFromBech32(v.OperatorAddress)
		// error should never happen as stored validator addresses must always be valid
		if err != nil {
			return err
		}
		// we just need to cast directly to sdk.AccAddress because the bech32 parsing
		// was previously validated and the underlying bytes are the same
		accAddr := sdk.AccAddress(operator.Bytes())
		r := k.GetValidatorRewards(ctx, accAddr)
		k.SetValidatorRewards(ctx, accAddr, validatorReward.Add(sdk.NewInt(r)).Int64())
	}
	return nil
}

// GetProportions gets the balance of the `MintedDenom` from minted coins
// and returns coins according to the `AllocationRatio`
func (k Keeper) GetProportions(ctx sdk.Context, mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, mintedCoin.Amount.ToDec().Mul(ratio).TruncateInt())
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetModuleAccountBalance gets the module account.
func (k Keeper) GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.AccountI {
	return k.accountKeeper.GetModuleAccount(ctx, moduleName)
}

func (k Keeper) FundCommunityPool(ctx sdk.Context) error {
	// If this account exists and has coins, fund the community pool.
	// The address hardcoded here is randomly generated with no keypair behind it. It will be empty and unused after the genesis file is applied.
	funder, err := sdk.AccAddressFromBech32("nois103y4f6h80lc45nr8chuzr3fyzqywm9n0d8fxzu")
	if err != nil {
		panic(err)
	}
	balances := k.bankKeeper.GetAllBalances(ctx, funder)
	if balances.IsZero() {
		return nil
	}
	return k.distrKeeper.FundCommunityPool(ctx, balances, funder)
}
