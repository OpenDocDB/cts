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
	"context"
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/FerretDB/wire"
	"github.com/FerretDB/wire/wirebson"
	"github.com/FerretDB/wire/wireclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/OpenDocDB/cts/opendocdb-cts/internal/data"
)

func TestConvertFixtures(t *testing.T) { //nolint:revive // exceeds number of lines for readability
	for name, tc := range map[string]struct {
		fixtures data.Fixtures
		expected string
	}{
		"simple": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument("_id", "int32", "v", int32(42)),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "int32", "v": 42}]);`,
		},
		"Binary": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument(
						"_id", "binary",
						"v", wirebson.Binary{
							Subtype: wirebson.BinaryUser,
							B:       []byte{42, 0, 13},
						},
					),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "binary", "v": BinData(128, "KgAN")}]);`,
		},
		"ObjectID": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument(
						"_id", "objectid",
						"v", wirebson.ObjectID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11},
					),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "objectid", "v": ObjectId("000102030405060708091011")}]);`,
		},
		"Bool": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument("_id", "bool-false", "v", false),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "bool-false", "v": false}]);`,
		},
		"Datetime": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument(
						"_id", "datetime",
						"v", time.Date(2021, 11, 1, 10, 18, 42, 123000000, time.UTC),
					),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "datetime", "v": ISODate("2021-11-01T10:18:42.123Z")}]);`,
		},
		"Null": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument("_id", "null", "v", wirebson.NullType{}),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "null", "v": null}]);`,
		},
		"Regex": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument(
						"_id", "regex",
						"v", wirebson.Regex{
							Pattern: "foo",
							Options: "i",
						},
					),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "regex", "v": RegExp("foo", "i")}]);`,
		},
		"Timestamp": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument(
						"_id", "timestamp",
						"v", wirebson.Timestamp(42<<32|13),
					),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "timestamp", "v": Timestamp({t: 42, i: 13})}]);`,
		},
		"Long": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument("_id", "int64", "v", int64(0)),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "int64", "v": Long(0)}]);`,
		},
		"Decimal": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument("_id", "decimal128", "v", wirebson.Decimal128{H: 42, L: 13}),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "decimal128", "v": Decimal128("42.13")}]);`,
		},
		"Infinity": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument("_id", "double-inf", "v", math.Inf(1)),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "double-inf", "v": Double(+Infinity)}]);`,
		},
		"NegativeInfinity": {
			fixtures: data.Fixtures{
				"c": []*wirebson.Document{
					wirebson.MustDocument("_id", "double-neg-inf", "v", math.Inf(-1)),
				},
			},
			expected: `
			db.c.insertMany([{"_id": "double-neg-inf", "v": Double(-Infinity)}]);`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual, err := ConvertFixtures(tc.fixtures)
			require.NoError(t, err)
			assert.Equal(t, unindent(tc.expected)+"\n", actual)

			f, err := os.CreateTemp("", name+".js")
			require.NoError(t, err)

			_, err = f.WriteString(actual)
			require.NoError(t, err)

			path, err := filepath.Abs(f.Name())
			require.NoError(t, err)

			require.NoError(t, f.Close())

			ctx := context.TODO()
			l := slog.Default()

			// cleanup database
			// TODO: use cli flag
			dbName := "test"

			collName := "c"

			conn, err := wireclient.Connect(ctx, "mongodb://127.0.0.1:27001/", l)
			require.NoError(t, err)

			t.Cleanup(func() {
				require.NoError(t, conn.Close())
				f.Close()
			})

			require.NoError(t, conn.Ping(ctx))

			_, _, err = conn.Request(ctx, wire.MustOpMsg(
				"dropDatabase", int32(1),
				"$db", dbName,
			))
			require.NoError(t, err)

			// run command against database

			err = runMongosh(l, path, os.Stderr)
			require.NoError(t, err)

			// fetch data from collection with wireclient from database

			_, body, err := conn.Request(ctx, wire.MustOpMsg(
				"find", collName,
				"$db", dbName,
			))
			require.NoError(t, err)

			doc, err := body.(*wire.OpMsg).DocumentDeep()
			require.NoError(t, err)

			// compare fetched data with inserted from fixtures

			// TODO fetch all of the cursor data
			assert.Equal(t, tc.fixtures["c"][0], doc.Get("cursor").(*wirebson.Document).Get("firstBatch").(*wirebson.Array).Get(0))

		})
	}
}

// TODO run in docker exec
func runMongosh(l *slog.Logger, path string, stdout io.Writer) error {
	cmd := exec.Command("mongosh", "--f", path, "mongodb://127.0.0.1:27001/")
	l.Debug(fmt.Sprintf("Running %s", strings.Join(cmd.Args, " ")))

	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed: %s", strings.Join(cmd.Args, " "), err)
	}

	return nil
}
