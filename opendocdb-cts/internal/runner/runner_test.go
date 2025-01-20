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

package runner

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/neilotoole/slogt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/OpenDocDB/cts/opendocdb-cts/internal/data"
)

func TestRunner(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	dir := filepath.Join("..", "..", "..", "cts")
	db := t.Name()

	f, tss, err := data.Load(dir, map[string]string{
		"Database": db,
	})
	require.NoError(t, err)

	r, err := New("mongodb://127.0.0.1:27001/"+db, slogt.New(t))
	require.NoError(t, err)

	err = r.Setup(ctx, f)
	require.NoError(t, err)

	for name, ts := range tss {
		err = r.Run(ctx, ts)
		assert.NoError(t, err, "%s", name)
	}
}
