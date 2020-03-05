package main

import (
	"fmt"
)

/*
func main() {
	fmt.Println("Hello!")
	port := 3000
	startWebServer(port, 5)
}

// func startWebServer(port int, numberOfRetries int) {
func startWebServer(port, numberOfRetries int) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}
*/

/*
func main() {
	port := 3000
	isStarted := startWebServer(port)
	fmt.Println(isStarted)
}

func startWebServer(port int) bool { // "bool" - returned data type
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	return true
}
*/

func main() {
	port := 3000
	// port, err := startWebServer(port)
	// fmt.Println(port, err)
	_, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	// return nil
	// return errors.New("Something went wrong")
	return port, nil
}
