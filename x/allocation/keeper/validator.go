package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

func (k Keeper) SetValidatorRewards(ctx sdk.Context, operator sdk.AccAddress, rewards types.ValidatorReward) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	// delete if rewards are zero
	if rewards.Rewards.IsZero() {
		prefixStore.Delete(operator)
		return
	}
	bz := k.cdc.MustMarshal(&rewards)
	prefixStore.Set(operator, bz)
}

func (k Keeper) GetValidatorRewards(ctx sdk.Context, operator sdk.AccAddress) types.ValidatorReward {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	if !prefixStore.Has(operator) {
		return types.ValidatorReward{}
	}
	bz := prefixStore.Get(operator)
	rewards := types.ValidatorReward{}
	k.cdc.MustUnmarshal(bz, &rewards)
	return rewards
}
