package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

func (k Keeper) SetValidatorRewards(ctx sdk.Context, operator sdk.AccAddress, rewards types.ValidatorAccumulatedRewards) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	bz := k.cdc.MustMarshal(&rewards)
	prefixStore.Set(operator, bz)
}

func (k Keeper) GetValidatorRewards(ctx sdk.Context, operator sdk.AccAddress) types.ValidatorAccumulatedRewards {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	if !prefixStore.Has(operator) {
		return types.ValidatorAccumulatedRewards{}
	}
	bz := prefixStore.Get(operator)
	rewards := types.ValidatorAccumulatedRewards{}
	k.cdc.MustUnmarshal(bz, &rewards)
	return rewards
}
