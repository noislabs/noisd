package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

// ClaimRewards claims the rewards for the given operator
func (k Keeper) ClaimRewards(ctx sdk.Context, operator sdk.AccAddress) (sdk.Coins, error) {
	rewardAmount := k.GetValidatorRewards(ctx, operator)
	if rewardAmount == 0 {
		return sdk.Coins{}, types.ErrNoRewards
	}
	// send the rewards to the operator
	reward := sdk.NewCoins(sdk.NewInt64Coin(k.stakingKeeper.BondDenom(ctx), rewardAmount))
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ValidatorRewardsPool, operator, reward)
	if err != nil {
		return sdk.Coins{}, err
	}
	// remove the rewards from the store
	k.DeleteValidatorRewards(ctx, operator)
	return reward, nil
}
