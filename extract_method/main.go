package main

import (
	"fmt"
	"strconv"
)

func assert(description string, expected string, actual string) {
	if expected != actual {
		panic(fmt.Sprintf("%s\nExpected:\n[%v] but was:\n[%v]", description, expected, actual))
	}
}

func main() {
	test := func(testFunc func(int,string) string) {
		assert("prints details with banner", "BANNER\nname:my_name\namount:123\n", testFunc(123, "my_name"))
	}

	fmt.Println("Before")
	test(Before)
	fmt.Println("After")
	test(After)
	fmt.Println("Excercise")
	testExcercise()
}

func testExcercise() {
	assert("when elements exist", "****Banner\nElement: a\nElement: b\nname:name\namount:1\ntotal:2\n", Excercise([]string{"a", "b"}, "name", 1))
	assert("when elements are empty", "****Banner\nname:name\namount:2\ntotal:0\n", Excercise([]string{}, "name", 2))
	assert("when total less than zero", "****Banner\nElement: a\nname:name\namount:-1\ntotal is less than zero\n", Excercise([]string{"a"}, "name", -1))
}

func Before(amount int, name string) string {
	printBanner := func() string{
		return "BANNER\n"
	}
	banner := printBanner()

	//print details
	banner += "name:" + name + "\n"
	banner += "amount:" + strconv.Itoa(amount) + "\n"

	return banner
}

func After(amount int, name string) string {
	printBanner := func() string{
		return "BANNER\n"
	}
	printDetails := func(amount int, name string) string {
		details := ""
		details += "name:" + name + "\n"
		details += "amount:" + strconv.Itoa(amount) + "\n"
		return details
	}
	banner := printBanner()
	details := printDetails(amount, name)

	return banner + details
}

func Excercise(elements []string, name string, amount int) string {
	owingStr := ""
	total := 0
	outstandingStr := ""

	//calculate outstanding
	for _, element := range elements {
		// print element
		outstandingStr += "Element: " + element + "\n"
	}

	//calculate total
	for range elements {
		total += amount
	}

	//
	// print details
	//

	//print banner
	owingStr += "****Banner" + "\n"

	// print elements
	owingStr += outstandingStr

	//print name
	owingStr += "name:" + name + "\n"

	//print amount
	owingStr += "amount:" + strconv.Itoa(amount) + "\n"

	// print total section
	if total<0 {
		//print alert message
		owingStr += "total is less than zero" + "\n"
	} else {
		// print total value
		owingStr += "total:" + strconv.Itoa(total) + "\n"

	}

	return owingStr
}
