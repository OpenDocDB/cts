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
		results  []TestSuiteResult
	}{
		"Empty": {
			results: []TestSuiteResult{},
			expected: "" +
				"Test Name     |Result\n" +
				"---------     |------\n",
		},
		"Simple": {
			results: []TestSuiteResult{
				{Name: "test1/foo", Passed: true},
				{Name: "test2/bar", Passed: false},
			},
			expected: "" +
				"Test Name     |Result\n" +
				"---------     |------\n" +
				"test1/foo     |✅\n" +
				"test2/bar     |❌\n",
		},
		"Sorted": {
			results: []TestSuiteResult{
				{Name: "foo", Passed: true},
				{Name: "bar", Passed: false},
			},
			expected: "" +
				"Test Name     |Result\n" +
				"---------     |------\n" +
				"bar           |❌\n" +
				"foo           |✅\n",
		},
		"LongName": {
			results: []TestSuiteResult{
				{Name: "LoremIpsumDolorSitAmetConsecteturAdipiscingElit", Passed: true},
				{Name: "b", Passed: false},
			},
			expected: "" +
				"Test Name                                           |Result\n" +
				"---------                                           |------\n" +
				"LoremIpsumDolorSitAmetConsecteturAdipiscingElit     |✅\n" +
				"b                                                   |❌\n",
		},
	}

	for Name, tc := range tests {
		t.Run(Name, func(t *testing.T) {
			result := ResultsTable(tc.results)
			assert.Equal(t, tc.expected, result)
		})
	}
}
