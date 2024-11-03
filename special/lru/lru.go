package main

import (
    "container/list"
)
// LRUCache структура кэша
type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
}

// Node представляет элемент кэша
type Node struct {
    key   int
    value int
}

// Конструктор LRUCache
func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[int]*list.Element),
        list:     list.New(),
    }
}

func (lru *LRUCache) Get(k int) int {
	if element, found := lru.cache[k]; found {
		lru.list.MoveToFront(element)
		return element.Value.(*Node).value
	}
	return -1
}

func (lru *LRUCache) Set(k, v int) {
	if element, found := lru.cache[k]; found {
		element.Value.(*Node).value = v
		lru.list.MoveToFront(element)
	} else {
		if lru.list.Len() == lru.capacity {
			back := lru.list.Back()
			if back != nil {
				lru.list.Remove(back)
			}
			delete(lru.cache, back.Value.(*Node).key)
		}
		newNode := &Node{key: k, value: v}
		newElement := lru.list.PushFront(newNode)
		lru.cache[k] = newElement
	}

}