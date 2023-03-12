package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/noislabs/noisd/x/allocation/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper

	paramstore paramtypes.Subspace
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, stakingKeeper types.StakingKeeper,
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
		paramstore: ps,
	}
}

// DistributeInflation distributes module-specific inflation
func (k Keeper) DistributeInflation(ctx sdk.Context) error {
	// Fee collector module account account contains newly minted coins and collected fees from transactions
	blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).GetAddress()
	blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, k.stakingKeeper.BondDenom(ctx))

	// if there is no inflation, return
	if blockInflation.IsZero() {
		return nil
	}
	params := k.GetParams(ctx)
	proportions := params.DistributionProportions
	// fund randomness rewards address
	randomnessRewardsCoin := k.GetProportions(ctx, blockInflation, proportions.RandomnessRewards)
	randomnessRewardsReceiver, err := sdk.AccAddressFromBech32(params.RandomnessRewardsReceiver)
	if err != nil {
		return err
	}
	if !randomnessRewardsCoin.IsZero() {
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, authtypes.FeeCollectorName, randomnessRewardsReceiver, sdk.NewCoins(randomnessRewardsCoin))
	}
	// fund validator rewards pool
	validatorRewardsCoins := k.GetProportions(ctx, blockInflation, proportions.ValidatorRewards)
	if !validatorRewardsCoins.IsZero() {
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, authtypes.FeeCollectorName, types.ValidatorRewardsPool, sdk.NewCoins(validatorRewardsCoins))
	}
	return nil
}

func (k Keeper) DistributeValidatorRewards(ctx sdk.Context) error {
	validators := k.stakingKeeper.GetLastValidators(ctx)
	validValidators := make([]stakingtypes.Validator, 0)
	for _, validator := range validators {
		// skip jailed and not bonded
		if validator.Jailed || validator.Status != stakingtypes.Bonded {
			continue
		}
		// TODO: retrieve self-delegation
		validValidators = append(validValidators, validator)
	}
	// TODO:  accumulate rewards
	fmt.Println(validValidators)
	return nil
}

// GetProportions gets the balance of the `MintedDenom` from minted coins
// and returns coins according to the `AllocationRatio`
func (k Keeper) GetProportions(ctx sdk.Context, mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, mintedCoin.Amount.ToDec().Mul(ratio).TruncateInt())
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
