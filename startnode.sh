#!/bin/bash
set -o errexit -o nounset -o pipefail
command -v shellcheck >/dev/null && shellcheck "$0"


# Create key under nickname "validator" first:
# noisd keys add validator
VALIDATOR_ADDR=$(noisd keys show validator -a)

# rm -rf $HOME/.noisd
rm -rf ~/.noisd/config/gentx
# setup chain
noisd init local-validator-001 --chain-id localnet-1 --overwrite
noisd prepare-genesis localnet-1

noisd add-genesis-account "$VALIDATOR_ADDR" 500000000000000unois
# add initial community pool balance of 10000 NOIS
noisd add-genesis-account nois103y4f6h80lc45nr8chuzr3fyzqywm9n0d8fxzu 10000000000unois
noisd gentx validator 1000000unois --chain-id localnet-1 --details "1"
noisd collect-gentxs
noisd validate-genesis
noisd tendermint reset-state
noisd tendermint unsafe-reset-all
