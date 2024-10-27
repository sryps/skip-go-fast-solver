//go:build tools

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	// development tools
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"

	// abigen for generating go bindings for evm abi
	_ "github.com/ethereum/go-ethereum/cmd/abigen"

	// mockery for generating mocks
	_ "github.com/vektra/mockery/v2"

	// protocol buffer tools
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/nefixestrada/protoc-gen-go-grpc-mock"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	// sql tools
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
