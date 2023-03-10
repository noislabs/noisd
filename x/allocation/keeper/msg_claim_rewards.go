package keeper

import (
	"context"

	"github.com/noislabs/noisd/x/allocation/types"
)

func (k msgServer) ClaimRewards(goCtx context.Context, req *types.MsgClaimRewards) (*types.MsgClaimRewardsResponse, error) {
	return nil, nil
}
