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

package mongosh

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/FerretDB/wire/wirebson"

	"github.com/OpenDocDB/cts/opendocdb-cts/internal/data"
)

// ConvertRequest converts wirebson's request Document to a mongosh's JavaScript `runCommand`.
func ConvertRequest(req *wirebson.Document) (string, error) {
	s, err := convert(req)
	if err != nil {
		return "", fmt.Errorf("mongosh.ConvertRequest: %w", err)
	}

	return "db.runCommand(" + s + ");\n", nil
}

// ConvertResponse converts wirebson's response Document to a mongosh's JavaScript `response` variable.
func ConvertResponse(res *wirebson.Document) (string, error) {
	s, err := convert(res)
	if err != nil {
		return "", fmt.Errorf("mongosh.ConvertResponse: %w", err)
	}

	return "response = " + s + "\n", nil
}

// ConvertTestSuites converts request and response from each testsuite
// to a mongosh's JavaScript format, and writes them to files in specified
// directory.
func ConvertTestSuites(testSuites data.TestSuites, outDir string) error {
	for tsName, ts := range testSuites {
		reqDir := filepath.Join(outDir, "requests", tsName)
		resDir := filepath.Join(outDir, "responses", tsName)

		if err := os.MkdirAll(reqDir, 0o777); err != nil {
			return err
		}

		if err := os.MkdirAll(resDir, 0o777); err != nil {
			return err
		}

		for tcName, tc := range ts {
			filename := fmt.Sprintf("%s.js", tcName)

			req, err := ConvertRequest(tc.Request)
			if err != nil {
				return err
			}

			res, err := ConvertResponse(tc.Response)
			if err != nil {
				return err
			}

			if err = os.WriteFile(filepath.Join(reqDir, filename), []byte(req), 0o666); err != nil {
				return err
			}

			if err = os.WriteFile(filepath.Join(resDir, filename), []byte(res), 0o666); err != nil {
				return err
			}
		}
	}

	return nil
}
