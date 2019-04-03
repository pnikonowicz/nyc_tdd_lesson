package main

import (
	"fmt"
	. "strings"
)

//TODO: Tests dont pass

//From Page 124 Introduce Explaining Variable
func Before(platform, browser string, resize int, wasInitialized bool) bool {
	if Compare(ToUpper(platform), "MAC") > -1 &&
		Compare(ToUpper(browser), "IE") > -1 &&
		wasInitialized && resize > 0 {
		return true
	}
	return false
}

func After(platform, browser string, resize int, wasInitialized bool) bool {
	isMacOs := Compare(ToUpper(platform), "MAC") > -1
	isIEBrowser := Compare(ToUpper(browser), "IE") > -1
	wasResized := resize > 0

	if isMacOs &&
		isIEBrowser &&
		wasInitialized &&
		wasResized {
		return true
	}
	return false
}

type Item struct {
	Quality int
	Name    string
	SellIn  int
}

func StudentExcercise(item Item) Item {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
		if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			if item.SellIn < 11 {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
			if item.SellIn < 6 {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
		}
	}
	return item
}

func assert(expected bool, actual bool) {
	if expected != actual {
		panic(fmt.Sprintf("Expected: [%v] but was: [%v]", expected, actual))
	}
}

func test(name string, testFunction func(string, string, int, bool) bool) {
	fmt.Println(fmt.Sprintf("Running Suite: %s", name))
	isMacOs := "MAC"
	isIeBrowser := "IE"
	wasResized := 1 // number > 0
	isInitialized := true

	notMacOs := "not mac"
	notIeBrowser := "not ie"
	notResized := -1 // number <= 0
	notInitialized := false

	assert(true, testFunction(isMacOs, isIeBrowser, wasResized, isInitialized))
	assert(false, testFunction(notMacOs, isIeBrowser, wasResized, isInitialized))
	assert(false, testFunction(isMacOs, notIeBrowser, wasResized, isInitialized))
	assert(false, testFunction(isMacOs, isIeBrowser, notResized, isInitialized))
	assert(false, testFunction(isMacOs, isIeBrowser, wasResized, notInitialized))
}

func main() {
	test("Before", Before)
	test("After", After)
}