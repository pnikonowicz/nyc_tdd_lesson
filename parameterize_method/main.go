package main

import "fmt"

func main() {
	assert := func (description string, expected float32, actual float32) {
		if expected != actual {
			panic(fmt.Sprintf("%s\nExpected:\n[%v] but was:\n[%v]", description, expected, actual))
		}
	}

	fmt.Println("BEFORE")
	{
		structUnderTest := Before{salary: 10}
		structUnderTest.tenPercentRaise()
		assert("tenPercentRaise", 11, structUnderTest.salary)

		structUnderTest = Before{salary: 10}
		structUnderTest.fivePercentRaise()
		assert("fivePercentRaise", 10.5, structUnderTest.salary)
	}

	fmt.Println("AFTER")
	{
		structUnderTest := After{salary: 10}
		structUnderTest.raise(1.1)
		assert("tenPercentRaise", 11, structUnderTest.salary)

		structUnderTest = After{salary: 10}
		structUnderTest.raise(1.05)
		assert("fivePercentRaise", 10.5, structUnderTest.salary)
	}

	fmt.Println("EXAMPLE")
	{
		assertExample := func (description string, expected int, actual int) {
			if expected != actual {
				panic(fmt.Sprintf("%s\nExpected:\n[%v] but was:\n[%v]", description, expected, actual))
			}
		}

		assertExample("when less than 10", 0, (&Example{lastUsage:0}).increment())
		assertExample("when greater than 10", 11, (&Example{lastUsage:10}).increment())
		assertExample("when greater than 100", 111, (&Example{lastUsage:100}).increment())

		assertExample("ExampleRefactored when less than 10", 0, ExampleRefactored(0))
		assertExample("ExampleRefactored when greater than 10", 11, ExampleRefactored(10))
		assertExample("ExampleRefactored when greater than 100", 111, ExampleRefactored(100))
	}
}

type Before struct {
	salary float32
}

func (b *Before) tenPercentRaise() {
	b.salary*=1.1
}

func (b *Before) fivePercentRaise() {
	b.salary*=1.05
}

type After struct {
	salary float32
}

func (a *After) raise(factor float32) {
	a.salary *= factor
}


type Example struct {
	lastUsage int
}

func (e *Example) increment() int {
	lastUsageAbove := func(x int) bool {
		return e.lastUsage >= x
	}

	if lastUsageAbove(10) {
		e.lastUsage += 1
	}

	if lastUsageAbove(100) {
		e.lastUsage += 10
	}

	return e.lastUsage
}

func ExampleRefactored(lastUsage int) int {
	if lastUsage >= 10 {
		lastUsage += 1
	}

	if lastUsage >= 100 {
		lastUsage += 10
	}

	return lastUsage
}