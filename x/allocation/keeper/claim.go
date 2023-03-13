package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

// ClaimRewards claims the rewards for the given operator
func (k Keeper) ClaimRewards(ctx sdk.Context, operator sdk.AccAddress) (sdk.Coins, error) {
	valRewards := k.GetValidatorRewards(ctx, operator)
	if valRewards.Rewards.IsZero() {
		return sdk.Coins{}, nil
	}
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ValidatorRewardsPool, operator, valRewards.Rewards)
	if err != nil {
		return sdk.Coins{}, err
	}
	// remove the rewards from the store
	k.DeleteValidatorRewards(ctx, operator)
	return valRewards.Rewards, nil
}
