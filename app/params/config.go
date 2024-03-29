package params

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

type WasmConfig struct {
	// SimulationGasLimit is the max gas that can be spent when executing a simulation TX.
	SimulationGasLimit uint64 `mapstructure:"simulation_gas_limit"`
	// QueryGasLimit is the maximum sdk gas (wasm and storage) that we allow for any x/wasm "smart" queries
	QueryGasLimit uint64 `mapstructure:"query_gas_limit"`
	// MemoryCacheSize defines the memory size for Wasm modules that we can keep cached to speed-up instantiation
	// The value is in MiB not bytes
	MemoryCacheSize uint64 `mapstructure:"memory_cache_size"`
}

// CustomAppConfig defines the configuration for the Nois app.
type CustomAppConfig struct {
	serverconfig.Config
	WASM WasmConfig `mapstructure:"wasm"`
}

const customAppTemplate = `
[wasm]
# This is the maximum sdk gas (wasm and storage) that we allow for any x/wasm "smart" queries
query_gas_limit = 4500000

# This the max gas that can be spent when executing a simulation TX
simulation_gas_limit = 4500000

# This defines the memory size for Wasm modules that we can keep cached to speed-up instantiation
# The value is in MiB not bytes
memory_cache_size = 512
`

func CustomConfigTempalte() string {
	return serverconfig.DefaultConfigTemplate + customAppTemplate
}

func DefaultConfig() (string, interface{}) {
	serverConfig := serverconfig.DefaultConfig()
	serverConfig.MinGasPrices = "0.05unois"
	customConfig := CustomAppConfig{
		Config: *serverConfig,
		WASM: WasmConfig{
			SimulationGasLimit: 4_500_000,
			QueryGasLimit:      4_500_000,
			MemoryCacheSize:    512,
		},
	}
	return CustomConfigTempalte(), customConfig
}
