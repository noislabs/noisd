package noistesting

import (
	"encoding/json"
	"time"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/std"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"
	app "github.com/noislabs/noisd/app"
	appparams "github.com/noislabs/noisd/app/params"
)

type EmptyOptions struct{}

func (EmptyOptions) Get(key string) interface{} {
	return nil
}

func NewApp(home string) *app.NoisApp {
	db := dbm.NewMemDB()
	encCdc := appparams.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encCdc.Amino)
	std.RegisterInterfaces(encCdc.InterfaceRegistry)
	app.ModuleBasics.RegisterLegacyAminoCodec(encCdc.Amino)
	app.ModuleBasics.RegisterInterfaces(encCdc.InterfaceRegistry)
	noisApp := app.NewNoisApp(
		log.NewNopLogger(),
		db,
		nil,              // no trace store
		true,             // load latest version
		map[int64]bool{}, // no skip upgrade heights
		home,
		5, // invariant check periods,
		encCdc,
		EmptyOptions{},
		nil, // empty wasm options
	)

	return noisApp
}

var defaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

func SetupNewApp(home string) *app.NoisApp {
	noisApp := NewApp(home)

	encCdc := appparams.MakeEncodingConfig()
	stateBytes, err := json.MarshalIndent(app.ModuleBasics.DefaultGenesis(encCdc.Codec), "", " ")
	if err != nil {
		panic(err)
	}

	// InitChain updates deliverState which is required when app.NewContext is called
	noisApp.InitChain(abci.RequestInitChain{
		ConsensusParams: defaultConsensusParams,
		AppStateBytes:   stateBytes,
	})

	return noisApp
}
