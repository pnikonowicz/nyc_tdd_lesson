package main

import (
	"fmt"
	. "strings"
)

//TODO: Student work example
//TODO: Tests dont pass


//For Page 124 Introduce Explaining Variable
func Before(platform, browser string, resize int, wasInitialized func() bool) bool {
	if Compare(ToUpper(platform), "MAC") > -1 &&
		Compare(ToUpper(browser), "IE") > -1 &&
		wasInitialized() && resize > 0 {
		return true
	}
	return false
}

func After(platform, browser string, resize int, wasInitialized func() bool) bool {
	isMacOs := Compare(ToUpper(platform), "MAC") > -1
	isIEBrowser := Compare(ToUpper(browser), "IE") > -1
	wasResized := resize > 0

	if isMacOs &&
		isIEBrowser &&
		wasInitialized() &&
		wasResized {
		return true
	}
	return false
}

func assert(expected bool, actual bool) {
	if expected != actual {
		panic(fmt.Sprintf("Expected: [%v] but was: [%v]", expected, actual))
	}
}

func test(testFunction func(string, string, int, func() bool) bool) {
	isMacOs := "MAC"
	isIeBrowser := "IE"
	wasResized := 1 // number > 0
	isInitialized := func() bool { return true }

	notMacOs := "not mac"
	notIeBrowser := "not ie"
	notResized := -1 // number <= 0
	notInitialized := func() bool { return false }

	assert(true, testFunction(isMacOs, isIeBrowser, wasResized, isInitialized))
	assert(false, testFunction(notMacOs, isIeBrowser, wasResized, isInitialized))
	assert(false, testFunction(isMacOs, notIeBrowser, wasResized, isInitialized))
	assert(false, testFunction(isMacOs, isIeBrowser, notResized, isInitialized))
	assert(false, testFunction(isMacOs, isIeBrowser, wasResized, notInitialized))
}

func main() {
	fmt.Println("Running Suite: Before")
	test(Before)

	fmt.Println("Running Suite: After")
	test(After)
}
