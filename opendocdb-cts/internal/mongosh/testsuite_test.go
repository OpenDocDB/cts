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
	"testing"

	"github.com/FerretDB/wire/wirebson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertRequest(t *testing.T) {
	req := wirebson.MustDocument(
		"find", "values",
		"filter", wirebson.MustDocument(
			"v", wirebson.MustDocument("$eq", int32(42)),
		),
		"sort", wirebson.MustDocument("_id", int32(1)),
		"$db", "test",
	)

	expected := `db.runCommand({"find": "values", "filter": {"v": {"$eq": Int32(42)}}, ` +
		`"sort": {"_id": Int32(1)}, "$db": "test"});` + "\n"

	actual, err := ConvertRequest(req)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
