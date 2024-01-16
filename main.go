package main

import (
	"fmt"
)

type LRUCache struct {
	Capacity   int
	CachedData KeyValuePairs
}

type KeyValuePair struct {
	Key   string
	Value interface{}
}

type KeyValuePairs []KeyValuePair

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
	}
}

func (c *LRUCache) Get(key string) interface{} {

	return c.search(key)
}

func (c *LRUCache) Set(key string) interface{} {

	return c.insert(key)
}

func (c *LRUCache) search(key string) interface{} {

	return -1
}

func (c *LRUCache) insert(key string) interface{} {

	return -1
}

func main() {
	fmt.Println("LRU Caching")
}
