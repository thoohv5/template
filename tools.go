//go:build tools
// +build tools

// Package tools manages development tool versions through the module system.
//
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/axw/gocov/gocov"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/matm/gocov-html"
	_ "golang.org/x/tools/cmd/goimports"
	_ "gotest.tools/gotestsum"

	_ "github.com/facebook/ent/entc/gen@v0.4.3"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/thoohv5/converter/cmd/converter"
	
	_ "github.com/swaggo/swag/cmd/swag@latest"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/files"
)
