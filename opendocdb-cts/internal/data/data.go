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

// Package data provides access to loading and saving CTS data.
package data

// Load parses JSON-encoded fixtures and test suites from the given directory (recursively).
func Load(dir string, vars map[string]string) (Fixtures, TestSuites, error) {
	fixtures, err := LoadFixtures(dir)
	if err != nil {
		return nil, nil, err
	}

	suites, err := LoadTestSuites(dir, vars)
	if err != nil {
		return nil, nil, err
	}

	return fixtures, suites, nil
}
