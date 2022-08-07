package lru_cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache[string, int](2)

	cache.Put("1", 1)
	cache.Put("2", 2)
	val, _ := cache.Get("1")
	if expected := 1; expected != val {
		t.Errorf("expect cached value to be [%d], got instead [%d]", expected, val)
	}

	cache.Put("3", 3) // evicts key 2
	if _, found := cache.Get("2"); found {
		t.Error("expect cached value to be blank")
	}

	cache.Put("4", 4) // evicts key 1
	if _, found := cache.Get("1"); found {
		t.Error("expect cached value to be blank")
	}

	val, _ = cache.Get("3")
	if expected := 3; expected != val {
		t.Errorf("expect cached value to be [%d], got instead [%d]", expected, val)
	}

	val, _ = cache.Get("4")
	if expected := 4; expected != val {
		t.Errorf("expect cached value to be [%d], got instead [%d]", expected, val)
	}
}
