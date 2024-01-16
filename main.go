package main

import (
	"fmt"
)

type LRUCache struct {
	Capacity   int
	CachedData map[string]*CachedItem
}

type CachedItem struct {
	Key      string
	Value    any
	previous *CachedItem
	next     *CachedItem
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity:   capacity,
		CachedData: make(map[string]*CachedItem),
	}
}

func (c *LRUCache) Get(key string) *CachedItem {
	if _, ok := c.CachedData[key]; !ok {
		return nil
	}

	return c.CachedData[key]
}

func (c *LRUCache) Set(key string) any {
	return c.insert(key)
}

func (c *LRUCache) insert(key string) any {
	return -1
}

func main() {
	fmt.Println("LRU Caching")
}
