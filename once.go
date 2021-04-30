package main

import "bytes"

type onceCounter struct {
	counter [][]byte
}

func (c *onceCounter) AssertOnce(key []byte) bool {
	for i := range c.counter {
		if bytes.Equal(c.counter[i], key) {
			return false
		}
	}
	c.counter = append(c.counter, key)
	return true
}
