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
		structUnderTest.tenPercentRaise()
		assert("tenPercentRaise", 11, structUnderTest.salary)

		structUnderTest = After{salary: 10}
		structUnderTest.fivePercentRaise()
	}

	fmt.Println("EXAMPLE")
	{
		assertExample := func (description string, expected int, actual int) {
			if expected != actual {
				panic(fmt.Sprintf("%s\nExpected:\n[%v] but was:\n[%v]", description, expected, actual))
			}
		}

		assertExample("when less than 10", 0, (&Example{lastUsage:0}).Example())
		assertExample("when greater than 10", 11, (&Example{lastUsage:10}).Example())
		assertExample("when greater than 100", 111, (&Example{lastUsage:100}).Example())

		assertExample("when less than 10", 0, (&Example{lastUsage:0}).ExampleRefactored())
		assertExample("when greater than 10", 11, (&Example{lastUsage:10}).ExampleRefactored())
		assertExample("when greater than 100", 111, (&Example{lastUsage:100}).ExampleRefactored())
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

func (a *After) tenPercentRaise() {
	a.raise(1.1)
}

func (a *After) fivePercentRaise() {
	a.raise(1.05)
}

type Example struct {
	lastUsage int
}

func (e *Example) Example() int {
	currentUsage := -1
	lastUsageAbove := func() bool {
		return e.lastUsage >= currentUsage
	}

	currentUsage = 10
	if lastUsageAbove() {
		e.lastUsage += 1
	}

	currentUsage = 100
	if lastUsageAbove() {
		e.lastUsage += 10
	}

	return e.lastUsage
}

func (e *Example) ExampleRefactored() int {
	isGreaterThanOrEqual := func(currentUsage int) bool {
		return e.lastUsage >= currentUsage
	}

	if isGreaterThanOrEqual(10) {
		e.lastUsage += 1
	}

	if isGreaterThanOrEqual(100) {
		e.lastUsage += 10
	}

	return e.lastUsage
}