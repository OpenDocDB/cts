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
	"errors"
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
		OutDir string `type:"path" arg:"" help:"Output directory."`
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
func convertCommand() error {
	fixtures, testSuites, err := data.Load(cli.Dir, nil)
	if err != nil {
		return err
	}

	fixtureDir := filepath.Join(cli.Convert.OutDir, "fixtures")

	if err = os.MkdirAll(fixtureDir, 0o766); err != nil {
		return err
	}

	for name, fx := range fixtures {
		var fxs string

		fxs, err = mongosh.ConvertFixtures(map[string]data.Fixture{name: fx})
		if err != nil {
			return err
		}

		if err = os.WriteFile(filepath.Join(fixtureDir, name+".js"), []byte(fxs), 0666); err != nil {
			return err
		}
	}

	for tsName, ts := range testSuites {
		reqDir := filepath.Join(cli.Convert.OutDir, "requests", tsName)
		resDir := filepath.Join(cli.Convert.OutDir, "responses", tsName)

		if err = os.MkdirAll(reqDir, 0o766); err != nil {
			return err
		}

		if err = os.MkdirAll(resDir, 0o766); err != nil {
			return err
		}

		for tcName, tc := range ts {
			filename := fmt.Sprintf("%s.js", tcName)

			req, res, err := convertTestCase(tc)
			if err != nil {
				return err
			}

			if err = os.WriteFile(filepath.Join(reqDir, filename), []byte(req), 0666); err != nil {
				return err
			}

			if err = os.WriteFile(filepath.Join(resDir, filename), []byte(res), 0666); err != nil {
				return err
			}
		}
	}

	return nil
}

// convertTestCase returns request and response from provided tc,
// converted to JavaScript format.
func convertTestCase(tc data.TestCase) (req string, res string, err error) {
	req, err = mongosh.ConvertRequest(tc.Request)
	if err != nil {
		return "", "", err
	}

	res, err = mongosh.ConvertResponse(tc.Response)
	if err != nil {
		return "", "", err
	}

	return req, res, err
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
		return errors.New("some tests failed")
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
		err = convertCommand()
	default:
		panic("unknown command")
	}

	kongCtx.FatalIfErrorf(err)
}
