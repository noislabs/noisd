package testing

import (
	"encoding/json"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"

	app "github.com/noislabs/noisd/app"
	appparams "github.com/noislabs/noisd/app/params"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

type EmptyOptions struct {
}

func (EmptyOptions) Get(key string) interface{} {
	return nil
}
func NewApp(home string) *app.NoisApp {
	db := dbm.NewMemDB()
	encCdc := appparams.MakeEncodingConfig()
	noisApp := app.NewNoisApp(
		log.NewNopLogger(),
		db,
		nil,              // no trace store
		true,             // load latest version
		map[int64]bool{}, // no skip upgrade heights
		home,
		5, //invariant check periods,
		encCdc,
		app.GetWasmEnabledProposals(),
		EmptyOptions{},
		nil, // empty wasm options
	)

	return noisApp
}

var defaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
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
	// Initialize the chain
	noisApp.InitChain(abci.RequestInitChain{})
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
	noisApp.Commit()
	return noisApp
}
