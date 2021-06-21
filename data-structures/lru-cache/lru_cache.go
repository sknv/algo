package lru_cache

import (
	"container/list"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRUCache struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element, capacity),
		queue:    list.New(),
	}
}

func (c *LRUCache) Get(key string) interface{} {
	element, found := c.items[key]
	if !found {
		return nil
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (c *LRUCache) Put(key string, value interface{}) {
	if element, found := c.items[key]; found {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return
	}

	if c.queue.Len() >= c.capacity {
		c.removeOldestItem()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}
	element := c.queue.PushFront(item)
	c.items[item.Key] = element
}

func (c *LRUCache) removeOldestItem() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}
