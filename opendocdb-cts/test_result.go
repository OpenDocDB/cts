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

package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"text/tabwriter"
)

// testResult represents the result of a test.
type testResult struct {
	name   string
	passed bool
}

// resultsTable returns the results of the tests in a formatted table.
// The table is sorted by test name.
func resultsTable(results []testResult) string {
	slices.SortFunc(results, func(a, b testResult) int { return cmp.Compare(a.name, b.name) })

	var sb strings.Builder
	w := tabwriter.NewWriter(&sb, 0, 0, 5, ' ', tabwriter.Debug)

	if _, err := fmt.Fprintln(w, "Test Name\tResult"); err != nil {
		panic(err)
	}

	if _, err := fmt.Fprintln(w, "---------\t------"); err != nil {
		panic(err)
	}

	for _, result := range results {
		status := "❌"
		if result.passed {
			status = "✅"
		}
		if _, err := fmt.Fprintf(w, "%s\t%s\n", result.name, status); err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	return sb.String()
}
