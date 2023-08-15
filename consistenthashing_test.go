package algorithm

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsistentHashAdd(t *testing.T) {
	hashing := NewConsistentHashing(0, nil)
	val, ok := hashing.Get("any")
	assert.False(t, ok)
	assert.Nil(t, val)
	for i := 0; i < 10; i++ {
		hashing.AddWithReplicas("node"+strconv.Itoa(i), 10)
	}
	keys := make(map[string]int)
	for i := 0; i < 1000; i++ {
		key, ok := hashing.Get("request" + strconv.Itoa(i))
		assert.True(t, ok)
		keys[key.(string)]++
	}
	mapData := make(map[string]int, len(keys))
	for k, v := range keys {
		mapData[k] = v
	}
	log.Println(mapData)
	// new node hash between node2 and node3
	hashing.AddWithReplicas("nodeNew", 1)
	keysNew := make(map[string]int)
	for i := 0; i < 1000; i++ {
		key, ok := hashing.Get("request" + strconv.Itoa(i))
		assert.True(t, ok)
		keysNew[key.(string)]++
	}
	mapDataAdd := make(map[string]int, len(keysNew))
	for k, v := range keysNew {
		mapDataAdd[k] = v
	}
	log.Println(mapDataAdd)
}

func TestConsistentHashRemove(t *testing.T) {
	hashing := NewConsistentHashing(0, nil)
	val, ok := hashing.Get("any")
	assert.False(t, ok)
	assert.Nil(t, val)
	for i := 0; i < 10; i++ {
		hashing.AddWithReplicas("node"+strconv.Itoa(i), 10)
	}
	hashing.AddWithReplicas("nodeNew", 1)
	keys := make(map[string]int)
	for i := 0; i < 1000; i++ {
		key, ok := hashing.Get("request" + strconv.Itoa(i))
		assert.True(t, ok)
		keys[key.(string)]++
	}
	mapDataAdd := make(map[string]int, len(keys))
	for k, v := range keys {
		mapDataAdd[k] = v
	}
	log.Println(mapDataAdd)
	keysNew := make(map[string]int)
	hashing.Remove("nodeNew")
	for i := 0; i < 1000; i++ {
		key, ok := hashing.Get("request" + strconv.Itoa(i))
		assert.True(t, ok)
		keysNew[key.(string)]++
	}
	mapRemove := make(map[string]int, len(keysNew))
	for k, v := range keysNew {
		mapRemove[k] = v
	}
	log.Println(mapRemove)
}
