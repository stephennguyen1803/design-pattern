package main

import "fmt"

type Lfu struct {
}

func (lfu *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy.")
}
