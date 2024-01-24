package main

import "fmt"

type Lru struct {
}

func (lru *Lru) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategry")
}
