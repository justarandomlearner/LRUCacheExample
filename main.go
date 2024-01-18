package main

import (
	"fmt"
	"log"
	"os"
)

var logr = log.New(os.Stderr, "ERROR:", 2)

type LRUCache struct {
	head       *CachedPair //leftmost item (last accessed item)
	tail       *CachedPair //rightmost item (least accesed item)
	capacity   int
	cachedData map[string]*CachedPair
}

type CachedPair struct {
	Key      string
	Value    any
	previous *CachedPair
	next     *CachedPair
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:   capacity,
		cachedData: make(map[string]*CachedPair),
	}
}

func (c *LRUCache) Get(key string) (any, error) {
	if c.head == nil {
		return nil, fmt.Errorf(("can't retrieve data from empty cache"))
	}

	cachedPair, ok := c.cachedData[key]
	if !ok {
		return nil, fmt.Errorf("no value stored for the given key")
	}

	// last acessed item must be at the leftmost position (head)
	c.setNewHead(c.cachedData[key])

	//if a cached item exists for the given key, it's value is not nil
	return cachedPair.Value, nil
}

func (c *LRUCache) setNewHead(newHead *CachedPair) {
	if c.head.next == nil {
		return
	}

	newHeadOldPrevious := newHead.previous
	newHeadOldNext := newHead.next

	newHeadOldPrevious.next = newHeadOldNext

	if newHeadOldNext != nil {
		newHeadOldNext.previous = newHeadOldPrevious
	}

	if newHead == c.tail {
		c.tail = newHeadOldPrevious
	}

	newHead.next = c.head
	newHead.previous = nil
	c.head = newHead

}

func (c *LRUCache) Insert(key string, value any) error {
	if c == nil {
		return fmt.Errorf("non-initialized cache")
	}

	if value == nil {
		return fmt.Errorf("cannot insert a nil value")
	}

	if c.head == nil {
		c.head = &CachedPair{Key: key, Value: value, previous: nil, next: nil}
		c.tail = c.head
		c.cachedData[key] = c.head
		return nil
	}

	if len(c.cachedData) == c.capacity {
		delKey := c.tail.Key
		c.tail = c.tail.previous
		c.tail.next = nil
		delete(c.cachedData, delKey)
	}

	oldTail := c.tail

	c.tail = &CachedPair{Key: key, Value: value}
	c.tail.previous = oldTail
	c.tail.next = nil

	oldTail.next = c.tail

	return nil
}

func main() {

	cache := NewLRUCache(3)

	if err := cache.Insert("oi", nil); err != nil {
		logr.Printf("%v", err)
	}
}
