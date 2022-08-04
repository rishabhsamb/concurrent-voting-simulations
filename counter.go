package main

import "sync"

type Counter struct {
	mu    *sync.RWMutex
	count uint64
}

func NewCounter() *Counter {
	return &Counter{&sync.RWMutex{}, 0}
}

func (c *Counter) add(num uint64) {
	(*c.mu).Lock()
	defer (*c.mu).Unlock()
	c.count += num
}

func (c *Counter) Inc() {
	c.add(1)
}

func (c *Counter) Get() uint64 {
	(*c.mu).RLock()
	defer (*c.mu).RUnlock()
	return c.count
}
