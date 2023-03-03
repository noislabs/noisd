#!/bin/bash
set -o errexit -o nounset -o pipefail
command -v shellcheck >/dev/null && shellcheck "$0"

# Use GNU version (Linux version) on Mac too. Otherwise the -i argument is incompatible.
gnused="$(command -v gsed || echo sed)"


# Create key under nickname "validator" first:
# noisd keys add validator
VALIDATOR_ADDR=$(noisd keys show validator -a)

# rm -rf $HOME/.noisd
rm -rf ~/.noisd/config/gentx
# setup chain
noisd init nois --chain-id localnet-1 --overwrite
"$gnused" -i 's/stake/unois/' ~/.noisd/config/genesis.json


noisd add-genesis-account "$VALIDATOR_ADDR" 500000000000000unois
noisd gentx validator 1000000unois --chain-id localnet-1 --details "1"
noisd collect-gentxs
noisd validate-genesis
noisd tendermint reset-state
noisd tendermint unsafe-reset-all


export CONFIG_DIR="$HOME/.noisd/config"

# Update app.toml
"$gnused" -i 's/minimum-gas-prices =.*$/minimum-gas-prices = "0.05unois"/' "$CONFIG_DIR/app.toml"

# Update block time settings (config.toml)
"$gnused" -i 's/^timeout_propose =.*$/timeout_propose = "2000ms"/' "$CONFIG_DIR/config.toml" \
  && "$gnused" -i 's/^timeout_propose_delta =.*$/timeout_propose_delta = "500ms"/' "$CONFIG_DIR/config.toml" \
  && "$gnused" -i 's/^timeout_prevote =.*$/timeout_prevote = "1s"/' "$CONFIG_DIR/config.toml" \
  && "$gnused" -i 's/^timeout_prevote_delta =.*$/timeout_prevote_delta = "500ms"/' "$CONFIG_DIR/config.toml" \
  && "$gnused" -i 's/^timeout_precommit =.*$/timeout_precommit = "1s"/' "$CONFIG_DIR/config.toml" \
  && "$gnused" -i 's/^timeout_precommit_delta =.*$/timeout_precommit_delta = "500ms"/' "$CONFIG_DIR/config.toml" \
  && "$gnused" -i 's/^timeout_commit =.*$/timeout_commit = "1750ms"/' "$CONFIG_DIR/config.toml"

