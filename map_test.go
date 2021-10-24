package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestMergeStrIFaceMaps(t *testing.T) {
	assert.Equal(t,
		map[string]interface{}{},
		MergeStrIFaceMaps(nil, nil),
	)

	assert.Equal(t,
		map[string]interface{}{"key1": 1, "key2": 2},
		MergeStrIFaceMaps(nil, map[string]interface{}{"key1": 1, "key2": 2}),
	)

	assert.Equal(t,
		map[string]interface{}{"key1": 1, "key2": 2},
		MergeStrIFaceMaps(map[string]interface{}{"key1": 1, "key2": 2}, nil),
	)

	assert.Equal(t,
		map[string]interface{}{"key1": 1, "key2": 2},
		MergeStrIFaceMaps(map[string]interface{}{"key1": 3, "key2": 4}, map[string]interface{}{"key1": 1, "key2": 2}),
	)

	assert.Equal(t,
		map[string]interface{}{"key1": 1, "key2": 2},
		MergeStrIFaceMaps(map[string]interface{}{"key1": 1, "key2": map[string]interface{}{"key3": "123"}}, map[string]interface{}{"key1": 1, "key2": 2}),
	)

	assert.Equal(t,
		map[string]interface{}{"key1": 1, "key2": map[string]interface{}{"key3": "456", "key4": 123, "key5": []int{1, 2, 3}}, "key3": 3},
		MergeStrIFaceMaps(
			map[string]interface{}{"key1": 1, "key2": map[string]interface{}{"key3": "123", "key5": []int{1, 2, 3}}},
			map[string]interface{}{"key1": 1, "key2": map[string]interface{}{"key3": "456", "key4": 123}, "key3": 3},
		),
	)
}
