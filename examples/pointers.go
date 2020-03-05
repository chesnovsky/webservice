package main

import "fmt"

func main() {
	// var firstName *string = new(string)
	// *firstName = "Woof-woof"
	// fmt.Println(*firstName)

	firstName := "Voiceh"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Vaclav"
	fmt.Println(ptr, *ptr)
}
