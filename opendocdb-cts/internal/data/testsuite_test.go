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

func TestLoadSaveTestSuites(t *testing.T) {
	dir := filepath.Join("..", "..", "..", "cts")
	vars := map[string]string{
		"Database": "qklck6njgh9o",
	}

	tss, err := LoadTestSuites(dir, vars)
	require.NoError(t, err)
	assert.Contains(t, tss, "query/eq", slices.Collect(maps.Keys(tss)))

	err = SaveTestSuites(tss, dir, vars)
	require.NoError(t, err)
}
