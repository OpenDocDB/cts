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
	"os"
	"path/filepath"
	"strings"

	"github.com/FerretDB/wire/wirebson"
)

// Fixture represents a collection of documents in a single collection.
type Fixture []*wirebson.Document

// Fixtures is a collection of named fixtures.
type Fixtures map[string]Fixture

// LoadFixture parses a JSON-encoded fixture from the given file.
func LoadFixture(file string) (Fixture, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("LoadFixture: %w", err)
	}

	d := json.NewDecoder(bytes.NewReader(b))
	d.DisallowUnknownFields()

	var f Fixture
	if err = d.Decode(&f); err != nil {
		return nil, fmt.Errorf("LoadFixture: %w", err)
	}

	return f, nil
}

// SaveFixture writes the JSON encoding of the fixture to the given file.
func SaveFixture(fixture Fixture, file string) error {
	b, err := json.MarshalIndent(fixture, "", "  ")
	if err != nil {
		return fmt.Errorf("SaveFixture: %w", err)
	}

	b = append(b, '\n')
	if err := os.WriteFile(file, b, 0o666); err != nil {
		return fmt.Errorf("SaveFixture: %w", err)
	}

	return nil
}

// LoadFixtures parses JSON-encoded fixtures from the given directory.
func LoadFixtures(dir string) (Fixtures, error) {
	matches, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("LoadFixtures: %w", err)
	}

	res := make(Fixtures, len(matches))

	for _, m := range matches {
		f, err := LoadFixture(m)
		if err != nil {
			return nil, fmt.Errorf("LoadFixtures: %w", err)
		}

		name := filepath.Base(m)
		res[strings.TrimSuffix(name, ".json")] = f
	}

	return res, nil
}

// SaveFixtures writes the JSON encoding of fixtures to the given directory.
func SaveFixtures(fixtures Fixtures, dir string) error {
	for name, f := range fixtures {
		if err := SaveFixture(f, filepath.Join(dir, name+".json")); err != nil {
			return fmt.Errorf("SaveFixtures: %w", err)
		}
	}

	return nil
}
