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

// Package runner provides CTS runner.
package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/url"
	"strings"

	"github.com/FerretDB/wire"
	"github.com/FerretDB/wire/wirebson"
	"github.com/FerretDB/wire/wireclient"
	"github.com/pmezard/go-difflib/difflib"

	"github.com/OpenDocDB/cts/opendocdb-cts/internal/data"
)

// Runner executes test suites.
type Runner struct {
	uri *url.URL
	l   *slog.Logger
	db  string
}

// New creates a new runner.
func New(uri string, l *slog.Logger) (*Runner, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if u.Path == "" || u.Path == "/" {
		return nil, fmt.Errorf("MongoDB URI's path can't be empty")
	}

	db := strings.TrimPrefix(u.Path, "/")
	u.Path = "/"

	return &Runner{
		uri: u,
		l:   l,
		db:  db,
	}, nil
}

// Setup sets up the test environment.
func (r *Runner) Setup(ctx context.Context, fixtures data.Fixtures) error {
	conn, err := wireclient.Connect(ctx, r.uri.String(), r.l)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err = conn.Ping(ctx); err != nil {
		return err
	}

	if _, _, err = conn.Request(ctx, wire.MustOpMsg("dropDatabase", int32(1), "$db", r.db)); err != nil {
		return err
	}

	for name, docs := range fixtures {
		documents := wirebson.MakeArray(len(docs))
		for _, doc := range docs {
			if err = documents.Add(doc); err != nil {
				return err
			}
		}

		rand.Shuffle(len(docs), func(i, j int) {
			d := documents.Get(i)
			if err == nil {
				err = documents.Replace(i, documents.Get(j))
			}
			if err == nil {
				err = documents.Replace(j, d)
			}
		})
		if err != nil {
			return err
		}

		_, body, err := conn.Request(ctx, wire.MustOpMsg(
			"insert", name,
			"documents", documents,
			"$db", r.db,
		))
		if err != nil {
			return err
		}

		d, err := body.(*wire.OpMsg).DecodeDeepDocument()
		if err != nil {
			return err
		}
		if d.Get("ok") != 1.0 {
			return fmt.Errorf("%s", d.LogMessage())
		}
	}

	return nil
}

// Run executes the given test suite.
func (r *Runner) Run(ctx context.Context, ts data.TestSuite) error {
	conn, err := wireclient.Connect(ctx, r.uri.String(), r.l)
	if err != nil {
		return err
	}
	defer conn.Close()

	var errs []error
	for name, tc := range ts {
		req, err := wire.NewOpMsg(tc.Request)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
			continue
		}

		_, body, err := conn.Request(ctx, req)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
			continue
		}

		resp, err := body.(*wire.OpMsg).DecodeDeepDocument()
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
			continue
		}

		expected, err := json.MarshalIndent(tc.Response, "", "  ")
		if err != nil {
			errs = append(errs, err)
			continue
		}

		actual, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if !bytes.Equal(expected, actual) {
			diff, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
				A:        difflib.SplitLines(string(expected)),
				FromFile: "expected",
				B:        difflib.SplitLines(string(actual)),
				ToFile:   "actual",
				Context:  1,
			})
			if err != nil {
				errs = append(errs, err)
				continue
			}

			err = fmt.Errorf(
				"%s: expected != actual\n\nExpected:\n%s\n\nActual:\n%s\n\nDiff:\n%s",
				name, expected, actual, diff,
			)
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

// RunGolden executes the given test suite and updates its file with actual results.
func (r *Runner) RunGolden(ctx context.Context, ts data.TestSuite, path string, vars map[string]string) error {
	conn, err := wireclient.Connect(ctx, r.uri.String(), r.l)
	if err != nil {
		return err
	}
	defer conn.Close()

	for name, tc := range ts {
		req, err := wire.NewOpMsg(tc.Request)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}

		_, body, err := conn.Request(ctx, req)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}

		resp, err := body.(*wire.OpMsg).DecodeDeepDocument()
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}

		tc.Response = resp
		ts[name] = tc

		if err = data.SaveTestSuite(ts, path, vars); err != nil {
			return err
		}
	}

	return nil
}
