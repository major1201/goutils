package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyMapSS(t *testing.T) {
	var map1 map[string]string // nil map
	map2 := map[string]string{
		"go":   "lang",
		"lang": "go",
	}

	assert.Nil(t, CopyMapSS(map1))
	assert.Equal(t, map2, CopyMapSS(map2))
}

func TestMergeMapSS(t *testing.T) {
	// case 1: normal
	assert.Equal(t,
		map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"},
		MergeMapSS(map[string]string{"key1": "value1"}, map[string]string{"key2": "value2"}, map[string]string{"key3": "value3"}),
	)

	// case2: nil base
	assert.Equal(t,
		map[string]string{"key2": "value2", "key3": "value3"},
		MergeMapSS(nil, map[string]string{"key2": "value2"}, map[string]string{"key3": "value3"}),
	)
	assert.Equal(t, map[string]string{}, MergeMapSS(nil, nil))
	assert.Nil(t, MergeMapSS(nil))

	// case 3: nil overrides
	assert.Equal(t,
		map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"},
		MergeMapSS(map[string]string{"key1": "value1"}, map[string]string{"key2": "value2"}, nil, map[string]string{"key3": "value3"}),
	)

	// case 4: override values
	assert.Equal(t,
		map[string]string{"key1": "value3"},
		MergeMapSS(map[string]string{"key1": "value1"}, map[string]string{"key1": "value2"}, map[string]string{"key1": "value3"}),
	)
}
