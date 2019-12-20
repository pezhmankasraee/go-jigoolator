package main

import "fmt"

func main() {
	sum := add(10, 20)
	fmt.Printf("SUM: %d\n", sum)

	fmt.Println("Hello ONE")
	fmt.Println("Hello TWO")
	fmt.Println("Hello, Three pezhmankasraeejob")
	fmt.Println("Hello, Four")
}

func add(a, b int) int {
	return a + b
}
