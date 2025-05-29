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
	"maps"
	"slices"
	"strings"

	"github.com/FerretDB/wire/wirebson"

	"github.com/OpenDocDB/cts/opendocdb-cts/internal/data"
)

// ConvertFixtures converts fixtures to a mongosh's JavaScript `insertMany` commands.
func ConvertFixtures(fixtures data.Fixtures) (string, error) {
	var buf strings.Builder

	for _, name := range slices.Sorted(maps.Keys(fixtures)) {
		f := fixtures[name]

		_, _ = fmt.Fprintf(&buf, "db.%s.insertMany(", name)

		arr := wirebson.MakeArray(len(f))
		for _, doc := range f {
			if err := arr.Add(doc); err != nil {
				return "", fmt.Errorf("mongosh.ConvertFixtures: %w", err)
			}
		}

		d, err := convert(arr)
		if err != nil {
			return "", fmt.Errorf("mongosh.ConvertFixtures: %w", err)
		}

		buf.WriteString(d)

		buf.WriteString(");\n")
	}

	return buf.String(), nil
}
