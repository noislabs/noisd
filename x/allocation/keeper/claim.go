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
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, operator, valRewards.Rewards)
	if err != nil {
		return sdk.Coins{}, err
	}
	k.SetValidatorRewards(ctx, operator, types.ValidatorAccumulatedRewards{})
	return valRewards.Rewards, nil
}

func (k Keeper) AllValidatorsRewards(ctx sdk.Context, operator sdk.AccAddress) ([]types.ValidatorAccumulatedRewards, error) {
	return nil, nil
}

func (k Keeper) SetValidatorsRewards(ctx sdk.Context, allRewards []types.ValidatorAccumulatedRewards) error {
	return nil
}
