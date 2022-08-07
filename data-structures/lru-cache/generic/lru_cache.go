package lru_cache

import (
	"container/list"
)

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

type LRUCache[K comparable, V any] struct {
	capacity int
	items    map[K]*list.Element
	queue    *list.List
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element, capacity),
		queue:    list.New(),
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	element, found := c.items[key]
	if !found {
		var zeroV V
		return zeroV, false
	}

	c.queue.MoveToFront(element)
	return element.Value.(*Item[K, V]).Value, true
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if element, found := c.items[key]; found {
		c.queue.MoveToFront(element)
		element.Value.(*Item[K, V]).Value = value
		return
	}

	if c.queue.Len() >= c.capacity {
		c.removeOldestItem()
	}

	item := &Item[K, V]{
		Key:   key,
		Value: value,
	}
	element := c.queue.PushFront(item)
	c.items[item.Key] = element
}

func (c *LRUCache[K, V]) removeOldestItem() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item[K, V])
		delete(c.items, item.Key)
	}
}
