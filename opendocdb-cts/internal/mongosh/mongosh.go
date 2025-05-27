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

// Package mongosh provides converts to mongosh's JavaScript syntax.
package mongosh

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/FerretDB/wire/wirebson"
)

// Converts wirebson value to their mongosh's JavaScript equivalent.
// The value should be fully decoded.
// [wirebson.RawArray] and [wirebson.DocumentArray] are not supported.
func convert(v any) (string, error) {
	switch v := v.(type) {
	// composite types

	case *wirebson.Document:
		var buf strings.Builder
		buf.WriteString("{\n")

		var comma bool
		for n, v := range v.All() {
			if comma {
				buf.WriteString(",\n")
			} else {
				comma = true
			}

			buf.WriteString(strconv.Quote(n))
			buf.WriteString(": ")

			s, err := convert(v)
			if err != nil {
				return "", fmt.Errorf("mongosh.convert: %w", err)
			}

			buf.WriteString(s)
		}

		buf.WriteString("\n}")
		return buf.String(), nil

	case *wirebson.Array:
		var buf strings.Builder
		buf.WriteString("[\n")

		for i, v := range v.All() {
			if i != 0 {
				buf.WriteString(",\n")
			}

			s, err := convert(v)
			if err != nil {
				return "", fmt.Errorf("mongosh.convert: %w", err)
			}

			buf.WriteString(s)
		}

		buf.WriteString("\n]")
		return buf.String(), nil

	// scalar types

	case float64:
		switch {
		case math.IsInf(v, -1):
			return `Double(-Infinity)`, nil
		case math.IsInf(v, 1):
			return `Double(+Infinity)`, nil
		default:
			res := strconv.FormatFloat(v, 'f', -1, 64)
			if !strings.Contains(res, ".") {
				// otherwise, it will be converted to int32 by mongosh
				res = "Double(" + res + ")"
			}
			return res, nil
		}
	case string:
		// For short strings, just return the quoted string
		if len(v) <= 80 {
			return fmt.Sprintf(`%q`, v), nil
		}
		
		// Split the string into parts using spaces where available
		parts := strings.SplitAfter(v, " ")
		var lines []string
		current := ""
		
		// Process each part
		for _, part := range parts {
			// If adding this part would exceed the line limit
			if len(current) + len(part) > 70 {
				// Save the current line and start a new one
				if len(current) > 0 {
					lines = append(lines, current)
					current = part
				} else {
					// Special case: a single part is too long
					// Split it to fit the line limit
					for len(part) > 70 {
						lines = append(lines, part[:70])
						part = part[70:]
					}
					current = part
				}
			} else {
				// Add part to the current line
				current += part
			}
		}
		
		// Add the last line if not empty
		if len(current) > 0 {
			lines = append(lines, current)
		}
		
		// Format the result with proper quotes and concatenation
		var result string
		for i, line := range lines {
			if i > 0 {
				result += " +\n"
			}
			result += fmt.Sprintf(`%q`, line)
		}
		
		return result, nil
	case wirebson.Binary:
		s := base64.RawStdEncoding.EncodeToString(v.B)
		return fmt.Sprintf(`BinData(%d, "%s")`, v.Subtype, s), nil
	case wirebson.ObjectID:
		return fmt.Sprintf(`ObjectId("%x")`, v), nil
	case bool:
		return fmt.Sprintf(`%t`, v), nil
	case time.Time:
		return fmt.Sprintf(`ISODate("%s")`, v.Format(time.RFC3339Nano)), nil
	case wirebson.NullType, nil:
		return `null`, nil
	case wirebson.Regex:
		return fmt.Sprintf(`RegExp("%s", "%s")`, v.Pattern, v.Options), nil
	case int32:
		// it will be stored as double by the legacy shell, but we don't care about it there
		return fmt.Sprintf(`%d`, v), nil
	case wirebson.Timestamp:
		return fmt.Sprintf(`Timestamp({t: %d, i: %d})`, v.T(), v.I()), nil
	case int64:
		// the legacy shell handles Long() / NumberLong() differently
		return fmt.Sprintf(`Long(%d)`, v), nil
	case wirebson.Decimal128:
		// TODO https://github.com/OpenDocDB/cts/issues/34
		return fmt.Sprintf(`Decimal128("%d.%d")`, v.H, v.L), nil
	default:
		return "", fmt.Errorf("mongosh.convert: invalid BSON type %T", v)
	}
}
