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
func convert(v any, buf *strings.Builder) error {
	switch v := v.(type) {
	// composite types

	case *wirebson.Document:
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

			err := convert(v, buf)
			if err != nil {
				return fmt.Errorf("mongosh.convert: %w", err)
			}
		}

		buf.WriteString("\n}")
		return nil

	case *wirebson.Array:
		buf.WriteString("[\n")

		for i, v := range v.All() {
			if i != 0 {
				buf.WriteString(",\n")
			}

			err := convert(v, buf)
			if err != nil {
				return fmt.Errorf("mongosh.convert: %w", err)
			}
		}

		buf.WriteString("\n]")
		return nil

	// scalar types

	case float64:
		switch {
		case math.IsInf(v, -1):
			buf.WriteString("Double(-Infinity)")
		case math.IsInf(v, 1):
			buf.WriteString("Double(+Infinity)")
		default:
			res := strconv.FormatFloat(v, 'f', -1, 64)
			if !strings.Contains(res, ".") {
				// otherwise, it will be converted to int32 by mongosh
				res = "Double(" + res + ")"
			}
			buf.WriteString(res)
		}
		return nil
	case string:
		words := strings.SplitAfter(v, " ")

		var sb strings.Builder

		var counter int
		for _, word := range words {
			if counter >= 80 {
				buf.WriteString(fmt.Sprintf(`%q`, sb.String()))
				sb.Reset()

				buf.WriteString(`" +\n"`)
				counter = 0
			}
			sb.WriteString(word)
			counter += len(word)
		}
		if sb.Len() != 0 {
			buf.WriteString(fmt.Sprintf(`%q`, sb.String()))
		}

		return nil
	case wirebson.Binary:
		s := base64.RawStdEncoding.EncodeToString(v.B)
		buf.WriteString(fmt.Sprintf("BinData(%d, \"%s\")", v.Subtype, s))
		return nil
	case wirebson.ObjectID:
		buf.WriteString(fmt.Sprintf("ObjectId(\"%x\")", v))
		return nil
	case bool:
		buf.WriteString(fmt.Sprintf("%t", v))
		return nil
	case time.Time:
		buf.WriteString(fmt.Sprintf("ISODate(\"%s\")", v.Format(time.RFC3339Nano)))
		return nil
	case wirebson.NullType, nil:
		buf.WriteString("null")
		return nil
	case wirebson.Regex:
		buf.WriteString(fmt.Sprintf("RegExp(\"%s\", \"%s\")", v.Pattern, v.Options))
		return nil
	case int32:
		// it will be stored as double by the legacy shell, but we don't care about it there
		buf.WriteString(fmt.Sprintf("%d", v))
		return nil
	case wirebson.Timestamp:
		buf.WriteString(fmt.Sprintf("Timestamp({t: %d, i: %d})", v.T(), v.I()))
		return nil
	case int64:
		// the legacy shell handles Long() / NumberLong() differently
		buf.WriteString(fmt.Sprintf("Long(%d)", v))
		return nil
	case wirebson.Decimal128:
		// TODO https://github.com/OpenDocDB/cts/issues/34
		buf.WriteString(fmt.Sprintf("Decimal128(\"%d.%d\")", v.H, v.L))
		return nil
	default:
		return fmt.Errorf("mongosh.convert: invalid BSON type %T", v)
	}
}
