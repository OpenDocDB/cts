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

package testresult

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultsTable(t *testing.T) {
	tests := map[string]struct {
		expected string
		results  []testCaseResult
	}{
		"Empty": {
			results: []testCaseResult{},
			expected: "" +
				"Test Name     |Result\n" +
				"---------     |------\n",
		},
		"Simple": {
			results: []testCaseResult{
				{name: "test1/foo", passed: true},
				{name: "test2/bar", passed: false},
			},
			expected: "" +
				"Test Name     |Result\n" +
				"---------     |------\n" +
				"test1/foo     |✅\n" +
				"test2/bar     |❌\n",
		},
		"Sorted": {
			results: []testCaseResult{
				{name: "foo", passed: true},
				{name: "bar", passed: false},
			},
			expected: "" +
				"Test Name     |Result\n" +
				"---------     |------\n" +
				"bar           |❌\n" +
				"foo           |✅\n",
		},
		"LongName": {
			results: []testCaseResult{
				{name: "LoremIpsumDolorSitAmetConsecteturAdipiscingElit", passed: true},
				{name: "b", passed: false},
			},
			expected: "" +
				"Test Name                                           |Result\n" +
				"---------                                           |------\n" +
				"LoremIpsumDolorSitAmetConsecteturAdipiscingElit     |✅\n" +
				"b                                                   |❌\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := resultsTable(tc.results)
			assert.Equal(t, tc.expected, result)
		})
	}
}
