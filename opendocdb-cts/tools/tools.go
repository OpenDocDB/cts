package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/quasilyte/go-consistent"
	_ "golang.org/x/tools/cmd/deadcode"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/stringer"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "mvdan.cc/gofumpt"
)

//go:generate go build -v -o ../../bin/ github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o ../../bin/ github.com/quasilyte/go-consistent
//go:generate go build -v -o ../../bin/ golang.org/x/tools/cmd/deadcode
//go:generate go build -v -o ../../bin/ golang.org/x/tools/cmd/goimports
//go:generate go build -v -o ../../bin/ golang.org/x/tools/cmd/stringer
//go:generate go build -v -o ../../bin/ golang.org/x/vuln/cmd/govulncheck
//go:generate go build -v -o ../../bin/ mvdan.cc/gofumpt
