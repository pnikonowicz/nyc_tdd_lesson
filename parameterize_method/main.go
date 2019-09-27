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
		structUnderTest := Begin{salary: 10}
		structUnderTest.tenPercentRaise()
		assert("tenPercentRaise", 11, structUnderTest.salary)

		structUnderTest = Begin{salary: 10}
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

		assertExample("when less than 10", 0, Example(0))
		assertExample("when greater than 10", 11, Example(10))
		assertExample("when greater than 100", 111, Example(100))

		assertExample("ExampleRefactored when less than 10", 0, ExampleRefactored(0))
		assertExample("ExampleRefactored when greater than 10", 11, ExampleRefactored(10))
		assertExample("ExampleRefactored when greater than 100", 111, ExampleRefactored(100))
	}
}

func Example(lastUsage int) int {
	if lastUsage >= 10 {
		lastUsage += 1
	}

	if lastUsage >= 100 {
		lastUsage += 10
	}

	return lastUsage
}

func ExampleRefactored(lastUsage int) int {
	conditionalAdd := func(cond int, amount int) {
		if lastUsage >= cond {
			lastUsage += amount
		}
	}

	conditionalAdd(10, 1)
	conditionalAdd(100, 10)

	return lastUsage
}

type Begin struct {
	salary float32
}

func (b *Begin) tenPercentRaise() {
	b.salary*=1.1
}

func (b *Begin) fivePercentRaise() {
	b.salary*=1.05
}

type After struct {
	salary float32
}

func (a *After) raise(factor float32) {
	a.salary *= factor
}