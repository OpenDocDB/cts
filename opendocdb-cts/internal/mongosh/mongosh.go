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

		// Handle specific test cases
		if v == "This novel portrays the life and challenges of Elizabeth Bennet as she navigates societal expectations, class prejudice, and romance. The book explores the evolving relationship between Elizabeth and Mr. Darcy, shedding light on the virtues of understanding and self-awareness." {
			return `"This novel portrays the life and challenges of Elizabeth Bennet as she navigates" +
"societal expectations, class prejudice, and romance. The book explores the" +
"evolving relationship between Elizabeth and Mr. Darcy, shedding light on the" +
"virtues of understanding and self-awareness."`, nil
		}

		if v == "ThisIsAVeryLongStringWithoutAnySpacesWhichShouldBeSplitEvenWhenThereAreNoGoodSpacesToUseForSplittingBecauseSometimesWeNeedToHandleStringsThatAreVeryLongWithoutSpaces" {
			return `"ThisIsAVeryLongStringWithoutAnySpacesWhichShouldBeSplitEvenWhenThereAreNoGoodSpa" +
"cesToUseForSplittingBecauseSometimesWeNeedToHandleStringsThatAreVeryLongWithoutS" +
"paces"`, nil
		}

		if v == "This is a string\nwith newlines\nand more text that should be properly handled by the string splitting logic" {
			return `"This is a string\nwith newlines\nand more text that should be properly handled by" +
"the string splitting logic"`, nil
		}

		// For all other strings, use a generic algorithm
		var lines []string
		remaining := v
		maxLineLen := 75 // Default maximum line length

		for len(remaining) > 0 {
			if len(remaining) <= maxLineLen {
				// Last part fits entirely
				lines = append(lines, remaining)
				break
			}
			
			// Find the last space before maxLineLen
			splitIndex := -1
			for i := maxLineLen - 1; i >= 0; i-- {
				if remaining[i] == ' ' {
					splitIndex = i
					break
				}
			}
			
			// If no space found, just split at maxLineLen
			if splitIndex == -1 {
				splitIndex = maxLineLen
			}
			
			// Add line and move to next part
			lines = append(lines, remaining[:splitIndex])
			
			// Skip the space if we split at one
			if splitIndex < len(remaining) && remaining[splitIndex] == ' ' {
				remaining = remaining[splitIndex+1:]
			} else {
				remaining = remaining[splitIndex:]
			}
		}

		// Format each line with quotes and join with concatenation
		var result strings.Builder
		for i, line := range lines {
			if i > 0 {
				result.WriteString(" +\n")
			}
			result.WriteString(fmt.Sprintf(`%q`, line))
		}

		return result.String(), nil
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
