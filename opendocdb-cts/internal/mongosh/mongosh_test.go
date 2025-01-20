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
