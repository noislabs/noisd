set -eux
#make install
VALIDATOR=$(noisd keys show validator -a)

# rm -rf $HOME/.noisd
rm -rf ~/.noisd/config/gentx
# setup chain
noisd init stargaze --chain-id localnet-1 --overwrite
sed -i 's/stake/unois/' ~/.noisd/config/genesis.json


noisd add-genesis-account $VALIDATOR 500000000000000unois
noisd gentx validator 1000000unois --chain-id localnet-1 --details "1"
noisd collect-gentxs
noisd validate-genesis
noisd tendermint reset-state
noisd tendermint unsafe-reset-all


export CONFIG_DIR="$HOME/.noisd/config"

# Update app.toml
sed -i 's/minimum-gas-prices =.*$/minimum-gas-prices = "0.05unois"/' $CONFIG_DIR/app.toml

# Update block time settings (config.toml)
sed -i 's/^timeout_propose =.*$/timeout_propose = "2000ms"/' $CONFIG_DIR/config.toml \
  && sed -i 's/^timeout_propose_delta =.*$/timeout_propose_delta = "500ms"/' $CONFIG_DIR/config.toml \
  && sed -i 's/^timeout_prevote =.*$/timeout_prevote = "1s"/' $CONFIG_DIR/config.toml \
  && sed -i 's/^timeout_prevote_delta =.*$/timeout_prevote_delta = "500ms"/' $CONFIG_DIR/config.toml \
  && sed -i 's/^timeout_precommit =.*$/timeout_precommit = "1s"/' $CONFIG_DIR/config.toml \
  && sed -i 's/^timeout_precommit_delta =.*$/timeout_precommit_delta = "500ms"/' $CONFIG_DIR/config.toml \
  && sed -i 's/^timeout_commit =.*$/timeout_commit = "1800ms"/' $CONFIG_DIR/config.toml
  
