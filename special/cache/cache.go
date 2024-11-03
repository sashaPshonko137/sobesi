package main

import "sync"

type Cache interface {
	Set(k, v string)
	Get(k string) (string, bool)
}

type cache struct {
	parts []map[string]string
	mu []sync.RWMutex
	n int
}

func New(n int) Cache {
	parts := make([]map[string]string, n)
	mu := make([]sync.RWMutex, n)
	for i := 0; i < n; i++ {
		parts[i] = make(map[string]string)
	}
	return &cache{
		parts: parts,
		mu: mu,
		n: n,
	}
}

func (c *cache) cacheFunc(k string) int {
	return int(k[0]) % c.n
}

func (c *cache) Set(k, v string) {
	n := c.cacheFunc(k)
	c.mu[n].Lock()
	defer c.mu[n].Unlock()
	c.parts[n][k] = v
}

func (c *cache) Get(k string) (string, bool) {
	n := c.cacheFunc(k)
	c.mu[n].RLock()
	defer c.mu[n].RUnlock()
	v, ok := c.parts[n][k]
	return v, ok
}
