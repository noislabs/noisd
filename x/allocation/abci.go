package allocation

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/keeper"
	"github.com/noislabs/noisd/x/allocation/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper, _ abci.RequestBeginBlock) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// for _, voteInfo := range req.LastCommitInfo.GetVotes() {
	// 	if voteInfo.SignedLastBlock {
	// 		voteInfo.Validator.GetAddress()
	// 	}
	// }

	if err := keeper.DistributeInflation(ctx); err != nil {
		panic(fmt.Sprintf("Error distribute inflation: %s", err.Error()))
	}
}
