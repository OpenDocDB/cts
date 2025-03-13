// Copyright 2025 Peter Farkas, Alexey Palazhchenko
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main provides OpenDocDB CTS tool.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"

	"github.com/OpenDocDB/cts/opendocdb-cts/internal/data"
	"github.com/OpenDocDB/cts/opendocdb-cts/internal/mongosh"
	"github.com/OpenDocDB/cts/opendocdb-cts/internal/runner"
)

// The cli struct represents all command-line commands, fields and flags.
// It's used for parsing the user input.
//
//nolint:vet // we don't care about fieldalignment here
var cli struct {
	Dir string `type:"path" default:"cts" help:"CTS directory."`

	Fmt struct{} `cmd:"" help:"Reformat CTS files."`

	Convert struct {
		OutDir string `arg:"" type:"path" help:"Output directory."`
	} `cmd:"" help:"Convert CTS files."`

	Run struct {
		URI    *url.URL `default:"mongodb://127.0.0.1:27017/cts" help:"Database URI."`
		Golden bool     `                                        help:"Update CTS files instead if failing tests."`
	} `cmd:"" help:"Run CTS."`
}

// fmtCommand implements the "fmt" command.
func fmtCommand() error {
	f, tss, err := data.Load(cli.Dir, nil)
	if err != nil {
		return err
	}

	if err = data.SaveFixtures(f, cli.Dir); err != nil {
		return err
	}

	return data.SaveTestSuites(tss, cli.Dir, nil)
}

// runCommand implements the "convert" command.
func convertCommand(ctx context.Context) error {
	f, err := data.LoadFixtures(cli.Dir)
	if err != nil {
		return err
	}

	b, err := mongosh.ConvertFixtures(f)

	return os.WriteFile(filepath.Join(cli.Convert.OutDir, "convert.json"), []byte(b), 0o666)
}

// runCommand implements the "run" command.
func runCommand(ctx context.Context, l *slog.Logger) error {
	r, err := runner.New(cli.Run.URI.String(), l)
	if err != nil {
		return err
	}

	vars := map[string]string{
		"Database": strings.TrimPrefix(cli.Run.URI.Path, "/"),
	}

	f, tss, err := data.Load(cli.Dir, vars)
	if err != nil {
		return err
	}

	if err = r.Setup(ctx, f); err != nil {
		return err
	}

	var failed bool
	for name, ts := range tss {
		if cli.Run.Golden {
			path := filepath.Join(cli.Dir, name+".json")
			err = r.RunGolden(ctx, ts, path, vars)
		} else {
			err = r.Run(ctx, ts)
		}

		if err == nil {
			l.InfoContext(ctx, name+": PASSED")
		} else {
			l.ErrorContext(ctx, name+": FAILED\n"+err.Error())
			failed = true
		}
	}

	if failed {
		return fmt.Errorf("some tests failed")
	}

	return nil
}

func main() {
	kongOptions := []kong.Option{
		kong.DefaultEnvars("CTS"),
	}

	ctx := context.Background()

	var err error

	kongCtx := kong.Parse(&cli, kongOptions...)
	switch kongCtx.Command() {
	case "fmt":
		err = fmtCommand()
	case "run":
		err = runCommand(ctx, slog.Default())
	case "convert <out-dir>":
		err = convertCommand(ctx)
	default:
		panic("unknown command")
	}

	kongCtx.FatalIfErrorf(err)
}
