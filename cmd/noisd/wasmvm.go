//go:build wasmvm
// +build wasmvm

package main

// This is a dummy file to make sure the wasmvm is imported as direct dependency
import (
	_ "github.com/CosmWasm/wasmvm"
)
