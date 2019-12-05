package main

import "fmt"

func X() string {
	return "A"
}

func main() {
	fmt.Printf(X())
}