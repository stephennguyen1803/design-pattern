package main

//Suppose you are building an In-Memory-Cache. Since it’s in memory, it has a limited size. Whenever it reaches its maximum size, some entries have to be evicted to free-up space. This can happen via several algorithms. Some of the popular algorithms are:
// 1-Least Recently Used (LRU): remove an entry that has been used least recently.
// 2-First In, First Out (FIFO): remove an entry that was created first.
// 3-Least Frequently Used (LFU): remove an entry that was least frequently used.

func main() {
	fifo := &Fifo{}
	cache := initCache(fifo)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	lfu := &Lfu{}
	cache.setEvictionAlgo(lfu)

	cache.add("e", "5")
}
