package main

import "fmt"

// const pi = 3.14
const (
	first = iota
	// second = 2 << iota
	second
)

const (
	third = iota
	fourth
)

func main() {
	fmt.Println(first, second, third, fourth)
}
