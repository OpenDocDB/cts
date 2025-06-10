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

	"github.com/OpenDocDB/cts/opendocdb-cts/internal/util/must"
)

// testCaseResult represents the result of a single test case.
type testCaseResult struct {
	name   string
	passed bool
}

// resultsTable returns the results of the test cases in a formatted table.
// The table is sorted by test case name.
func resultsTable(results []testCaseResult) string {
	slices.SortFunc(results, func(a, b testCaseResult) int { return cmp.Compare(a.name, b.name) })

	var sb strings.Builder
	w := tabwriter.NewWriter(&sb, 0, 0, 5, ' ', tabwriter.Debug)

	_ = must.NotFail(fmt.Fprintln(w, "Test Case Name\tResult"))
	_ = must.NotFail(fmt.Fprintln(w, "---------\t------"))

	for _, result := range results {
		status := "❌"
		if result.passed {
			status = "✅"
		}

		_ = must.NotFail(fmt.Fprintf(w, "%s\t%s\n", result.name, status))
	}

	must.NoError(w.Flush())

	return sb.String()
}
