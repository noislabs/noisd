package app

import (
	"encoding/json"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

type GenesisState map[string]json.RawMessage

func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	genesis := ModuleBasics.DefaultGenesis(cdc)
	wasmGen := wasmtypes.GenesisState{
		Params: wasmtypes.Params{
			CodeUploadAccess:             wasmtypes.AllowNobody,
			InstantiateDefaultPermission: wasmtypes.AccessTypeEverybody,
		},
	}
	genesis[wasmtypes.ModuleName] = cdc.MustMarshalJSON(&wasmGen)
	return genesis
}
