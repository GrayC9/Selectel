#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

echo "=== Building loglint ==="
cd "$ROOT_DIR"
go build -o "$SCRIPT_DIR/loglint" ./cmd/loglint/

echo ""
echo "=== Running loglint on example/main.go ==="
echo ""

cd "$SCRIPT_DIR"
./loglint ./... 2>&1 || true

echo ""
echo "=== Done ==="

rm -f "$SCRIPT_DIR/loglint"