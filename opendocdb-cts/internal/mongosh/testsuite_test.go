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

	expected := `
	db.runCommand({
	"find": "values",
	"filter": {
	"v": {
	"$eq": 42
	}
	},
	"sort": {
	"_id": 1
	}
	});`

	actual, err := ConvertRequest(req)
	require.NoError(t, err)
	assert.Equal(t, unindent(expected)+"\n", actual)
}

func TestConvertResponse(t *testing.T) {
	res := wirebson.MustDocument(
		"cursor", wirebson.MustDocument(
			"firstBatch", wirebson.MustArray(
				wirebson.MustDocument("_id", "int32-zero", "v", int32(0)),
				wirebson.MustDocument("_id", "int64-zero", "v", int64(0)),
			),
			"id", int64(0),
			"ns", "test.values",
		),
	)

	expected := `
	response = {
	"cursor": {
	"firstBatch": [
	{
	"_id": "int32-zero",
	"v": 0
	},
	{
	"_id": "int64-zero",
	"v": Long(0)
	}
	],
	"id": Long(0),
	"ns": "test.values"
	}
	}`

	actual, err := ConvertResponse(res)
	require.NoError(t, err)
	assert.Equal(t, unindent(expected)+"\n", actual)
}

func TestConvertResponseLongString(t *testing.T) {
	longText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in " +
		"reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, " +
		"sunt in culpa qui officia deserunt mollit anim id est laborum."

	res := wirebson.MustDocument(
		"cursor", wirebson.MustDocument(
			"firstBatch", wirebson.MustArray(
				wirebson.MustDocument("_id", "long-string", "v", longText),
				wirebson.MustDocument("_id", "int64-zero", "v", int64(0)),
			),
			"id", int64(0),
			"ns", "test.values",
		),
	)

	expected := `
	response = {
	"cursor": {
	"firstBatch": [
	{
	"_id": "long-string",
	"v": Lorem
	},
	],
	"id": Long(0),
	"ns": "test.values"
	}
	}`

	actual, err := ConvertResponse(res)
	require.NoError(t, err)
	assert.Equal(t, unindent(expected)+"\n", actual)
}
