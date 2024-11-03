package main

import (
	"fmt"
	"sync"
)

// Set представляет собой множество с синхронизацией
type Set[T comparable] struct { // Используем встроенное ограничение comparable
	m  map[T]struct{}
	mu sync.RWMutex
}

// New создает новый Set
func New[T comparable]() Set[T] {
	return Set[T]{
		m: make(map[T]struct{}),
	}
}

// Has проверяет наличие элемента в Set
func (s *Set[T]) Has(k T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.m[k]
	return ok
}

// Set добавляет элемент в Set
func (s *Set[T]) Set(k T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = struct{}{}
}

// Intersection находит пересечение двух множеств
func (s *Set[T]) Intersection(o *Set[T]) map[T]struct{} {
	res := make(map[T]struct{})
	s.mu.RLock()
	defer s.mu.RUnlock()
	o.mu.RLock()
	defer o.mu.RUnlock()

	for k := range s.m {
		if o.Has(k) {
			res[k] = struct{}{}
		}
	}
	return res
}

func main() {
	a := New[int]()
	b := New[int]()

	a.Set(1)
	a.Set(2)
	a.Set(3)

	b.Set(2)
	b.Set(3)
	b.Set(4)

	intersection := a.Intersection(&b)
	fmt.Println(intersection) // Вывод: map[2:{} 3:{}]
}
