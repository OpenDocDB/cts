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

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/FerretDB/wire/wirebson"
)

// TestCase represents a single request / expected response pair.
type TestCase struct {
	Request  *wirebson.Document `json:"request"`
	Response *wirebson.Document `json:"response"`
}

// TestSuite is an ordered set of named test cases.
// They are sorted by their names.
// Test cases may depend on previous ones.
type TestSuite map[string]TestCase

// TestSuites is an unordered set of named suites.
// They should be independent from each other.
type TestSuites map[string]TestSuite

// LoadTestSuite parses a JSON-encoded test suite from the given file.
func LoadTestSuite(file string, vars map[string]string) (TestSuite, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("LoadTestSuite: %w", err)
	}

	var r io.Reader = bytes.NewReader(b)

	if vars != nil {
		t, err := template.New("").Option("missingkey=error").Parse(string(b))
		if err != nil {
			return nil, fmt.Errorf("LoadTestSuite: %s: %w", file, err)
		}

		var buf bytes.Buffer
		if err = t.Execute(&buf, vars); err != nil {
			return nil, fmt.Errorf("LoadTestSuite: %s: %w", file, err)
		}

		r = &buf
	}

	d := json.NewDecoder(r)
	d.DisallowUnknownFields()

	var res TestSuite
	if err := d.Decode(&res); err != nil {
		return nil, fmt.Errorf("LoadTestSuite: %s: %w", file, err)
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("LoadTestSuite: %s: empty test suite", file)
	}

	for name, tc := range res {
		if tc.Request == nil {
			return nil, fmt.Errorf("LoadTestSuite: %s: %s: missing request", file, name)
		}

		if tc.Response == nil {
			return nil, fmt.Errorf("LoadTestSuite: %s: %s: missing response", file, name)
		}
	}

	return res, nil
}

// SaveTestSuite writes the JSON encoding of the test suite to the given file.
func SaveTestSuite(ts TestSuite, file string, vars map[string]string) error {
	b, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return fmt.Errorf("SaveTestSuite: %s: %w", file, err)
	}

	for k, v := range vars {
		b = bytes.ReplaceAll(b, []byte(v), []byte("{{."+k+"}}"))
	}

	b = append(b, '\n')
	if err = os.WriteFile(file, b, 0o666); err != nil {
		return fmt.Errorf("SaveTestSuite: %w", err)
	}

	return nil
}

// LoadTestSuites parses JSON-encoded test suites from the given directory (recursively).
func LoadTestSuites(dir string, vars map[string]string) (TestSuites, error) {
	var paths []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".json") || strings.HasSuffix(path, ".fixture.json") {
			return nil
		}

		paths = append(paths, path)
		return nil
	})
	if err != nil {
		return TestSuites{}, err
	}

	res := make(TestSuites, len(paths))
	for _, p := range paths {
		ts, err := LoadTestSuite(p, vars)
		if err != nil {
			return nil, fmt.Errorf("LoadTestSuites: %w", err)
		}

		name, err := filepath.Rel(dir, strings.TrimSuffix(p, ".json"))
		if err != nil {
			return nil, fmt.Errorf("LoadTestSuites: %w", err)
		}

		res[name] = ts
	}

	return res, nil
}

// SaveTestSuites writes the JSON encoding of test suites to the given directory.
func SaveTestSuites(tss TestSuites, dir string, vars map[string]string) error {
	for name, ts := range tss {
		if err := SaveTestSuite(ts, filepath.Join(dir, name+".json"), vars); err != nil {
			return err
		}
	}

	return nil
}
