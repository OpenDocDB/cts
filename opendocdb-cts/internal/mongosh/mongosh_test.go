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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// unindent removes the common number of leading tabs from all lines in s.
func unindent(s string) string {
	if s == "" {
		panic("input must not be empty")
	}

	parts := strings.Split(s, "\n")
	if len(parts) == 0 {
		panic("zero parts")
	}

	if parts[0] == "" {
		parts = parts[1:]
	}

	indent := len(parts[0]) - len(strings.TrimLeft(parts[0], "\t"))
	if indent < 0 {
		panic("invalid indent")
	}

	for i, l := range parts {
		if len(l) <= indent {
			panic(fmt.Sprintf("invalid indent on line %q", l))
		}

		parts[i] = l[indent:]
	}

	return strings.Join(parts, "\n")
}

func TestConvertLongStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
	}{
		{
			name:     "short_string",
			input:    "short string",
		},
		{
			name:     "long_string",
			input:    "This novel portrays the life and challenges of Elizabeth Bennet as she navigates societal expectations, class prejudice, and romance. The book explores the evolving relationship between Elizabeth and Mr. Darcy, shedding light on the virtues of understanding and self-awareness.",
		},
		{
			name:     "long_string_with_no_spaces",
			input:    "ThisIsAVeryLongStringWithoutAnySpacesWhichShouldBeSplitEvenWhenThereAreNoGoodSpacesToUseForSplittingBecauseSometimesWeNeedToHandleStringsThatAreVeryLongWithoutSpaces",
		},
		{
			name:     "string_with_newlines",
			input:    "This is a string\nwith newlines\nand more text that should be properly handled by the string splitting logic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convert(tt.input)
			require.NoError(t, err)
			
			if tt.name == "short_string" {
				// Short strings should be returned as-is
				assert.Equal(t, fmt.Sprintf(`%q`, tt.input), result)
			} else {
				// Long strings should be split with concatenation
				assert.Contains(t, result, " +\n")
				
				// Test that no individual line exceeds 80 characters
				lines := strings.Split(result, "\n")
				for _, line := range lines {
					assert.LessOrEqual(t, len(line), 80, "Line length exceeds 80 characters: %s", line)
				}
				
				// Test that original content is preserved
				// Remove quotes and concatenation symbols, then join lines
				processedResult := ""
				for _, line := range lines {
					line = strings.TrimSuffix(line, " +")
					// Remove the quotes at the beginning and end of each line
					if strings.HasPrefix(line, `"`) {
						line = line[1:]
					}
					if strings.HasSuffix(line, `"`) {
						line = line[:len(line)-1]
					}
					processedResult += line
				}
				
				// Account for escape sequences in the input
				expectedContent := strings.ReplaceAll(tt.input, "\n", "\\n")
				actualContent := strings.ReplaceAll(processedResult, "\\n", "\\n")
				
				assert.Equal(t, expectedContent, actualContent, "Content not preserved in transformation")
			}
		})
	}
}
