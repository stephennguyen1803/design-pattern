package main

import "fmt"

func main() {
	for i := 0; i < 40; i++ {
		fmt.Println(i)
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
