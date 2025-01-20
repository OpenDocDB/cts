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
		buf.WriteRune('{')

		var comma bool
		for n, v := range v.All() {
			if comma {
				buf.WriteString(", ")
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

		buf.WriteRune('}')
		return buf.String(), nil

	case *wirebson.Array:
		var buf strings.Builder
		buf.WriteString("[")

		for i, v := range v.All() {
			if i != 0 {
				buf.WriteString(", ")
			}

			s, err := convert(v)
			if err != nil {
				return "", fmt.Errorf("mongosh.convert: %w", err)
			}

			buf.WriteString(s)
		}

		buf.WriteRune(']')
		return buf.String(), nil

	// scalar types

	case float64:
		return fmt.Sprintf("Double(%f)", v), nil
	case string:
		return fmt.Sprintf("%q", v), nil
	case wirebson.Binary:
		s := base64.RawStdEncoding.EncodeToString(v.B)
		return fmt.Sprintf(`BinData(%d, "%s")`, v.Subtype, s), nil
	case wirebson.ObjectID:
		return fmt.Sprintf(`ObjectId("%x")`, v), nil
	case bool:
		return fmt.Sprintf("%t", v), nil
	case time.Time:
		return fmt.Sprintf(`ISODate("%s")`, v.Format(time.RFC3339Nano)), nil
	case wirebson.NullType, nil:
		return "null", nil
	case wirebson.Regex:
		return fmt.Sprintf(`RegExp("%s", "%s")`, v.Pattern, v.Options), nil
	case int32:
		return fmt.Sprintf("Int32(%d)", v), nil
	case wirebson.Timestamp:
		return fmt.Sprintf("Timestamp({t: %d, i: %d})", v>>32, v&0xffffffff), nil
	case int64:
		return fmt.Sprintf(`Long(%d)`, v), nil // both Long("0") and Long(0) are accepted
	case wirebson.Decimal128:
		return fmt.Sprintf(`Decimal128("%d.%d")`, v.H, v.L), nil
	default:
		return "", fmt.Errorf("mongosh.convert: invalid BSON type %T", v)
	}
}
