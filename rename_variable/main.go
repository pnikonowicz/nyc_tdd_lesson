package main

import "fmt"

func main() {
	assert := func(description string, expected string, actual string) {
		if expected != actual {
			panic(fmt.Sprintf("%s \nexpected: %s but was: %s", description, expected, actual))
		}
	}

	fmt.Println("BEFORE")
	{
		assert("equals expected string", "HI", Before())
	}

	fmt.Println("AFTER")
	{
		assert("equals expected string", "HI", After())
	}

	fmt.Println("EXAMPLE")
	{
		assert("equals expected string", "ABC", Example())
	}
}

func Before() string {
	something := "HI"
	return something
}

func After() string {
	somethingElse := "HI"
	return somethingElse
}

func Example() string {
	a := "A"
	b := "B"
	c := "C"

	return a+b+c
}
