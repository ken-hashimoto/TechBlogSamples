#!/bin/bash

echo "=== Comparing golang.org/x/sync versions ==="
echo

echo "1. Testing with v0.13.0 (traditional behavior):"
echo "-------------------------------------------"
sed -i '' 's/v0.14.0/v0.13.0/' go.mod
go mod tidy > /dev/null 2>&1
echo "Running: go run 1_panic_propagation/main.go"
go run 1_panic_propagation/main.go
echo "Exit code: $?"
echo

echo "2. Testing with v0.14.0 (new panic propagation):"
echo "-----------------------------------------------"
sed -i '' 's/v0.13.0/v0.14.0/' go.mod
go mod tidy > /dev/null 2>&1
echo "Running: go run 1_panic_propagation/main.go"
go run 1_panic_propagation/main.go
echo "Exit code: $?"
echo

echo "=== Comparison complete ===" 