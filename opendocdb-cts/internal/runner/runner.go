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
	"maps"
	"math/rand/v2"
	"net/url"
	"slices"
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
		return nil, errors.New("MongoDB URI's path can't be empty")
	}

	db := strings.TrimPrefix(u.Path, "/")
	u.Path = "/"

	return &Runner{
		uri: u,
		l:   l,
		db:  db,
	}, nil
}

// Setup sets up the test environment, recreating the database and fixtures.
func (r *Runner) Setup(ctx context.Context, fixtures data.Fixtures) error {
	// FIXME refactor to keep connection
	conn := wireclient.ConnectPing(ctx, r.uri.String(), r.l)
	if conn == nil {
		return fmt.Errorf("failed to connect to %s", r.uri.String())
	}

	defer conn.Close()

	username := r.uri.User.Username()
	password, _ := r.uri.User.Password()
	if username != "" {
		if err := conn.Login(ctx, username, password, ""); err != nil {
			return err
		}
	}

	_, _, err := conn.Request(ctx, wire.MustOpMsg(
		"dropDatabase", int32(1),
		"$db", r.db,
	))
	if err != nil {
		return err
	}

	for name, docs := range fixtures {
		documents := wirebson.MakeArray(len(docs))
		for _, doc := range docs {
			if err = documents.Add(doc); err != nil {
				return err
			}
		}

		rand.Shuffle(documents.Len(), func(i, j int) {
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

		d, err := body.(*wire.OpMsg).DocumentDeep()
		if err != nil {
			return err
		}
		if d.Get("ok") != 1.0 {
			return fmt.Errorf("%s", d.LogMessage())
		}
	}

	return nil
}

// prepareDocument returns document with multiple changes applied recursively:
//   - empty or duplicate field names are disallowed;
//   - fields are sorted by name.
//
// TODO https://github.com/OpenDocDB/cts/issues/89
func prepareDocument(doc *wirebson.Document) (*wirebson.Document, error) {
	res := wirebson.MakeDocument(doc.Len())

	names := slices.Compact(slices.Sorted(doc.Fields()))
	if doc.Len() != len(names) {
		return nil, errors.New("document contains duplicate fields")
	}

	for _, n := range names {
		if n == "" {
			return nil, errors.New("empty field names are not allowed")
		}

		v := doc.Get(n)
		if v == nil {
			return nil, errors.New("document is missing a field")
		}

		var fv any
		var err error

		switch v := v.(type) {
		case *wirebson.Document:
			fv, err = prepareDocument(v)
		case *wirebson.Array:
			fv, err = prepareArray(v)
		default:
			fv = v
		}
		if err != nil {
			return nil, err
		}

		if err := res.Add(n, fv); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// prepareArray returns array with multiple changes applied recursively.
// This function does not change anything, but calls [prepareDocument] for nested arrays.
func prepareArray(arr *wirebson.Array) (*wirebson.Array, error) {
	res := wirebson.MakeArray(arr.Len())

	for v := range arr.Values() {
		var fv any
		var err error

		switch v := v.(type) {
		case *wirebson.Document:
			fv, err = prepareDocument(v)
		case *wirebson.Array:
			fv, err = prepareArray(v)
		default:
			fv = v
		}

		if err != nil {
			return nil, err
		}
		if err = res.Add(fv); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// runTestCase runs a single test cases with fixes.
func (*Runner) runTestCase(ctx context.Context, conn *wireclient.Conn, tc data.TestCase) error {
	resp, err := prepareDocument(tc.Response)
	if err != nil {
		return err
	}

	expected, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	req, err := wire.NewOpMsg(tc.Request)
	if err != nil {
		return err
	}

	_, body, err := conn.Request(ctx, req)
	if err != nil {
		return err
	}

	resp, err = body.(*wire.OpMsg).DocumentDeep()
	if err != nil {
		return err
	}

	resp, err = prepareDocument(resp)
	if err != nil {
		return err
	}

	actual, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
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
			return err
		}

		return fmt.Errorf(
			"expected != actual\n\nExpected:\n%s\n\nActual:\n%s\n\nDiff:\n%s",
			expected, actual, diff,
		)
	}

	return nil
}

// Run executes the given test suite.
// Test cases are sorted by their names.
func (r *Runner) Run(ctx context.Context, ts data.TestSuite) error {
	// FIXME refactor to keep connection
	conn := wireclient.ConnectPing(ctx, r.uri.String(), r.l)
	if conn == nil {
		return fmt.Errorf("failed to connect to %s", r.uri.String())
	}

	defer conn.Close()

	username := r.uri.User.Username()
	password, _ := r.uri.User.Password()
	if username != "" {
		if err := conn.Login(ctx, username, password, ""); err != nil {
			return err
		}
	}

	var errs []error
	for _, name := range slices.Sorted(maps.Keys(ts)) {
		tc := ts[name]

		if err := r.runTestCase(ctx, conn, tc); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
		}
	}

	return errors.Join(errs...)
}

// RunGolden executes the given test suite and updates its file with actual results.
// Test cases are sorted by their names.
func (r *Runner) RunGolden(ctx context.Context, ts data.TestSuite, path string, vars map[string]string) error {
	// FIXME refactor to keep connection
	conn := wireclient.ConnectPing(ctx, r.uri.String(), r.l)
	if conn == nil {
		return fmt.Errorf("failed to connect to %s", r.uri.String())
	}

	defer conn.Close()

	username := r.uri.User.Username()
	password, _ := r.uri.User.Password()
	if username != "" {
		if err := conn.Login(ctx, username, password, ""); err != nil {
			return err
		}
	}

	for _, name := range slices.Sorted(maps.Keys(ts)) {
		tc := ts[name]

		req, err := wire.NewOpMsg(tc.Request)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}

		_, body, err := conn.Request(ctx, req)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}

		resp, err := body.(*wire.OpMsg).DocumentDeep()
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

func init() {
	wire.Debug = true
}
