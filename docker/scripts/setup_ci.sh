#!/bin/sh
set -o errexit -o nounset
command -v shellcheck >/dev/null && shellcheck "$0"

PASSWORD=${PASSWORD:-1234567890}
CHAIN_ID=${CHAIN_ID:-nois-testing}
MONIKER=${MONIKER:-nois-moniker}

TOKEN=${TOKEN:-unois}

# both types of tokens
START_BALANCE="1000000000$TOKEN"

echo "Creating genesis ..."
rm -rf "${HOME}/.noisd"
noisd init --chain-id "$CHAIN_ID" "$MONIKER"
noisd prepare-genesis "$CHAIN_ID"

cd "${HOME}/.noisd"

# this is essential for sub-1s block times (or header times go crazy)
sed -i 's/"time_iota_ms": "1000"/"time_iota_ms": "10"/' config/genesis.json

echo "Setting up validator ..."
if ! noisd keys show validator 2>/dev/null; then
  echo "Validator does not yet exist. Creating it ..."
  (
    echo "$PASSWORD"
    echo "$PASSWORD"
  ) | noisd keys add validator
fi
# hardcode the validator account for this instance
echo "$PASSWORD" | noisd add-genesis-account validator "$START_BALANCE"

echo "Setting up accounts ..."
# (optionally) add a few more genesis accounts
for addr in "$@"; do
  echo "$addr"
  noisd add-genesis-account "$addr" "$START_BALANCE"
done

echo "Creating genesis tx ..."
SELF_DELEGATION="3000000$TOKEN" # 3 NOIS (leads to a voting power of 3)
(
  echo "$PASSWORD"
  echo "$PASSWORD"
  echo "$PASSWORD"
) | noisd gentx validator "$SELF_DELEGATION" --offline --chain-id "$CHAIN_ID" --moniker="$MONIKER"

ls -lA config/gentx/

echo "Collecting genesis tx ..."
noisd collect-gentxs

# so weird, but found I needed the -M flag after lots of debugging odd error messages
# happening when redirecting stdout
jq -S -M . < config/genesis.json > genesis.tmp
mv genesis.tmp config/genesis.json
chmod a+rx config/genesis.json

# Custom settings in config.toml
sed -i"" \
  -e 's/^timeout_propose =.*$/timeout_propose = "100ms"/' \
  -e 's/^timeout_propose_delta =.*$/timeout_propose_delta = "100ms"/' \
  -e 's/^timeout_prevote =.*$/timeout_prevote = "100ms"/' \
  -e 's/^timeout_prevote_delta =.*$/timeout_prevote_delta = "100ms"/' \
  -e 's/^timeout_precommit =.*$/timeout_precommit = "100ms"/' \
  -e 's/^timeout_precommit_delta =.*$/timeout_precommit_delta = "100ms"/' \
  -e 's/^timeout_commit =.*$/timeout_commit = "200ms"/' \
  "config/config.toml"
