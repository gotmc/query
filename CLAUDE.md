# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Single-package Go library (`github.com/gotmc/query`) providing typed convenience functions for parsing string responses from a `Querier` interface. Designed for instrument communication (e.g., SCPI commands) where query results come back as strings that need conversion to Go types (bool, int, float64, string). Each type has a direct version and an `f` (format) variant that accepts `fmt.Sprintf`-style arguments. All functions take a `context.Context` as their first parameter.

## Commands

```bash
# Format, vet, and test
just check

# Lint (requires staticcheck)
just lint

# Test coverage HTML report
just cover

# Run a single test
go test -run TestBool ./...
```

Both `Makefile` and `Justfile` exist; the Justfile is more current (uses `GOEXPERIMENT=loopvar`).

## Architecture

- **`query.go`** — All library code: `Querier` interface and typed query functions (`Bool`, `Int`, `Float64`, `String` + `f` variants). `Int` handles scientific notation strings by falling back to `ParseFloat`.
- **`query_test.go`** — Table-driven tests using a mock `query` struct that implements `Querier` with a `map[string]string`.
- No external dependencies (stdlib only).
