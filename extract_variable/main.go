package main

import (
	"fmt"
	. "strings"
)

//From Page 124 Introduce Explaining Variable
func Before(platform, browser string, resize int, wasInitialized bool) bool {
	if ToUpper(platform) == "MAC" &&
		ToUpper(browser) == "IE" &&
		wasInitialized && resize > 0 {
		return true
	}
	return false
}

func After(platform, browser string, resize int, wasInitialized bool) bool {
	isMacOs := ToUpper(platform) == "MAC"
	isIEBrowser := ToUpper(browser) == "IE"
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

func Excercise(item Item) Item {
	if item.Quality < 50 {
		if item.Name == "backstage" {
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

func assert(description string, expected bool, actual bool) {
	if expected != actual {
		panic(fmt.Sprintf("%s\nExpected: [%v] but was: [%v]", description, expected, actual))
	}
}

func test(name string, testFunction func(string, string, int, bool) bool) {
	fmt.Println(fmt.Sprintf("Running Suite: %s", name))
	isMacOs := "MAC"
	isIeBrowser := "IE"
	wasResized := 1 // number > 0
	isInitialized := true

	notMacOs := "windows"
	notIeBrowser := "chrome"
	notResized := -1 // number <= 0
	notInitialized := false

	assert("success case", true, testFunction(isMacOs, isIeBrowser, wasResized, isInitialized))
	assert("false if not mac", false, testFunction(notMacOs, isIeBrowser, wasResized, isInitialized))
	assert("false if not ie", false, testFunction(isMacOs, notIeBrowser, wasResized, isInitialized))
	assert("false if not resize", false, testFunction(isMacOs, isIeBrowser, notResized, isInitialized))
	assert("false if not initialized", false, testFunction(isMacOs, isIeBrowser, wasResized, notInitialized))
}

func testExcercise() {
	fmt.Println("Suite: Excercise")
	assertEquals := func(description string, expected Item, input Item) {
		actual := Excercise(input)
		fail := func(check string) {
			panic(fmt.Sprintf("%s: %s\nexpected: %v actual: %v", check, description, expected, actual))
		}

		if expected.Quality != actual.Quality {
			fail("Quality")
		}

		if expected.Name != actual.Name {
			fail("Name")
		}

		if expected.SellIn != actual.SellIn {
			fail("SellIn")
		}
	}

	itemQualityLessThanFifty := func(i Item) Item { i.Quality = 3; return i }
	itemQualityGreaterThanFifty := func(i Item) Item { i.Quality = 52; return i }
	itemNameNotBackstage := func(i Item) Item { i.Name = "not backstage"; return i }
	itemNameIsBackstage := func(i Item) Item { i.Name = "backstage"; return i }
	itemSellinLessThanSix := func(i Item) Item { i.SellIn = 2; return i }
	itemSellinBetweenSixAndEleven := func(i Item) Item { i.SellIn = 8; return i }
	itemSellinEleven := func(i Item) Item { i.SellIn = 11; return i }
	itemLargeSellIn := func(i Item) Item { i.SellIn = 20; return i }
	incrementQualityTwice := func(i Item) Item { i.Quality = i.Quality + 2; return i }
	incrementQualityOnce := func(i Item) Item { i.Quality = i.Quality + 1; return i }

	item := itemQualityGreaterThanFifty(Item{})
	assertEquals("does nothing if item quality greater than fifty", item, item)

	item = itemNameNotBackstage(itemQualityGreaterThanFifty(Item{}))
	assertEquals("does nothing if not backstage",item,  item)

	item = itemLargeSellIn(itemLargeSellIn(itemNameIsBackstage(itemQualityLessThanFifty(Item{}))))
	assertEquals("does nothing if backstage but sell in too large", item, item)

	item = itemSellinEleven(itemNameIsBackstage(itemQualityLessThanFifty(Item{})))
	assertEquals("does nothing if item sellin is exactly 11", item, item)

	item = itemSellinLessThanSix(itemNameIsBackstage(itemQualityLessThanFifty(Item{})))
	assertEquals("increments quality twice", incrementQualityTwice(item), item)

	item = itemSellinBetweenSixAndEleven(itemNameIsBackstage(itemQualityLessThanFifty(Item{})))
	assertEquals("increments quality once", incrementQualityOnce(item), item)
}

func main() {
	test("Before", Before)
	test("After", After)
	testExcercise()
}
