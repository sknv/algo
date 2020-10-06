package lru_cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("1", 1)
	cache.Put("2", 2)
	if expected, val := 1, cache.Get("1"); expected != val {
		t.Errorf("expect cached value to be [%d], got instead [%d]", expected, val)
	}

	cache.Put("3", 3) // evicts key 2
	if val := cache.Get("2"); val != nil {
		t.Error("expect cached value to be blank")
	}

	cache.Put("4", 4) // evicts key 1
	if val := cache.Get("1"); val != nil {
		t.Error("expect cached value to be blank")
	}

	if expected, val := 3, cache.Get("3"); expected != val {
		t.Errorf("expect cached value to be [%d], got instead [%d]", expected, val)
	}

	if expected, val := 4, cache.Get("4"); expected != val {
		t.Errorf("expect cached value to be [%d], got instead [%d]", expected, val)
	}
}
