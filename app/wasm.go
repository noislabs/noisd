package app

import (
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

var (

	// WasmProposalsEnabled, if set to "true", enables all proposals for x/wasm.
	WasmProposalsEnabled = "true"

	// EnableSpecificWasmProposals, if set, must be comma-separated list of values
	// that are all a subset of "EnableAllProposals", which takes precedence over
	// WasmProposalsEnabled.
	//
	// See: https://github.com/CosmWasm/wasmd/blob/main/x/wasm/types/proposal.go#L34-L47
	EnableSpecificWasmProposals = ""
)

// GetWasmEnabledProposals parses the WasmProposalsEnabled and
// EnableSpecificWasmProposals values to produce a list of enabled proposals to
// pass into the application.
func GetWasmEnabledProposals() []wasm.ProposalType {
	if EnableSpecificWasmProposals == "" {
		if WasmProposalsEnabled == "true" {
			return wasm.EnableAllProposals
		}

		return wasm.DisableAllProposals
	}

	chunks := strings.Split(EnableSpecificWasmProposals, ",")

	proposals, err := wasm.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}

	return proposals
}

// NoisGasRegister creates instance with default values
func NoisGasRegister() wasmkeeper.WasmGasRegister {
	register := wasmkeeper.DefaultGasRegisterConfig()
	register.CompileCost = 3 // gas per byte
	return wasmkeeper.NewWasmGasRegister(register)
}

func NoisGasRegisterOption() wasm.Option {
	return wasmkeeper.WithGasRegister(NoisGasRegister())
}
