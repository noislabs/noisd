package allocation

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/keeper"
	"github.com/noislabs/noisd/x/allocation/types"
)

// InitGenesis initializes the alloc module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context,
	k keeper.Keeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	genState types.GenesisState,
) {
	validatorRewardsPool := k.GetModuleAccount(ctx, types.ValidatorRewardsPool)
	k.GetModuleAccount(ctx, types.ModuleName)
	k.SetParams(ctx, genState.Params)
	err := k.FundCommunityPool(ctx)
	if err != nil {
		panic(err)
	}
	bondDenom := stakingKeeper.BondDenom(ctx)
	totalRewards := sdk.NewInt(0)
	for _, reward := range genState.ValidatorRewards {
		acc, err := sdk.AccAddressFromBech32(reward.Address)
		if err != nil {
			panic(fmt.Sprintf("allocation genesis: %s is not a valid address", reward.Address))
		}
		amount := reward.Rewards.AmountOf(bondDenom)
		totalRewards = totalRewards.Add(amount)
		k.SetValidatorRewards(ctx, acc, amount.Int64())
	}
	totalBalance := bankKeeper.GetBalance(ctx, validatorRewardsPool.GetAddress(), bondDenom)
	if !totalBalance.Amount.Equal(totalRewards) {
		panic(fmt.Sprintf("allocation genesis: total rewards %s does not match total balance %s", totalRewards.String(), totalBalance.Amount.String()))
	}
}

// ExportGenesis returns the alloc module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:           k.GetParams(ctx),
		ValidatorRewards: k.ValidatorRewards(ctx),
	}
}
