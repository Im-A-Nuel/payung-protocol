#!/usr/bin/env bash
# Regenerates the go-ethereum bindings in internal/chain from the current
# Foundry build output. Run after any change to PayungPool.sol or IDRP.sol.
set -euo pipefail

BACKEND_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONTRACTS_DIR="$BACKEND_DIR/../contracts"
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

(cd "$CONTRACTS_DIR" && forge build >/dev/null)

for contract in PayungPool IDRP; do
    (cd "$CONTRACTS_DIR" && forge inspect "$contract" abi --json) > "$TMP_DIR/$contract.abi.json"
    abigen \
        --abi "$TMP_DIR/$contract.abi.json" \
        --pkg chain \
        --type "$contract" \
        --out "$BACKEND_DIR/internal/chain/$(echo "$contract" | tr '[:upper:]' '[:lower:]').go"
    echo "generated internal/chain/$(echo "$contract" | tr '[:upper:]' '[:lower:]').go"
done
