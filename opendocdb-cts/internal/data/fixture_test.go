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

package data

import (
	"maps"
	"path/filepath"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadSaveFixtures(t *testing.T) {
	dir := filepath.Join("..", "..", "..", "cts")

	f, err := LoadFixtures(dir)
	require.NoError(t, err)

	require.Contains(t, f, "values", slices.Collect(maps.Keys(f)))
	assert.NotEmpty(t, f["values"])

	err = SaveFixtures(f, dir)
	require.NoError(t, err)
}
