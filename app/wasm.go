package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// NoisGasRegister creates instance with default values
func NoisGasRegister() wasmtypes.GasRegister {
	register := wasmtypes.DefaultGasRegisterConfig()
	register.CompileCost = 3 // gas per byte
	return wasmtypes.NewWasmGasRegister(register)
}

func NoisGasRegisterOption() wasmkeeper.Option {
	return wasmkeeper.WithGasRegister(NoisGasRegister())
}
