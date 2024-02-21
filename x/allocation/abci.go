package allocation

import (
	"fmt"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/keeper"
	"github.com/noislabs/noisd/x/allocation/types"
)

func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper, _ abci.RequestBeginBlock) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	if err := keeper.DistributeInflation(ctx); err != nil {
		panic(fmt.Sprintf("error distribute inflation: %s", err.Error()))
	}
}
