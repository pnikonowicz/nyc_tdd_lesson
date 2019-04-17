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
		item.Quality = item.Quality + 1
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

	assert("1", true, testFunction(isMacOs, isIeBrowser, wasResized, isInitialized))
	assert("2", false, testFunction(notMacOs, isIeBrowser, wasResized, isInitialized))
	assert("3", false, testFunction(isMacOs, notIeBrowser, wasResized, isInitialized))
	assert("4", false, testFunction(isMacOs, isIeBrowser, notResized, isInitialized))
	assert("5", false, testFunction(isMacOs, isIeBrowser, wasResized, notInitialized))
}

func excercise() {
	fmt.Println("Suite: Excercise")
	assertEquals := func(description string, expected Item, input Item) {
		actual := Excercise(input)
		fail := func() {panic(fmt.Sprintf("%s\nexpected: %v actual: %v", description, expected, actual))}

		if expected.Quality != actual.Quality {
			fail()
		}

		if expected.Name != actual.Name {
			fail()
		}

		if expected.SellIn != actual.SellIn {
			fail()
		}
	}
	
	itemQualityLessThanFifty := func(i Item) Item {i.Quality= 3; return i}
	itemQualityGreaterThanFifty := func(i Item) Item {i.Quality= 52; return i}
	itemNameNotBackstage := func(i Item) Item {i.Name = "not backstage"; return i}
	itemNameIsBackstage := func(i Item) Item {i.Name = "backstage"; return i}
	itemSellinLessThanSix := func(i Item) Item {i.SellIn = 2; return i}
	itemSellinBetweenSixAndEleven := func(i Item) Item {i.SellIn = 8; return i}
	itemLargeSellIn := func(i Item) Item {i.SellIn = 20; return i}
	incrementQualityTwice := func(i Item) Item {i.Quality = i.Quality+2; return i}
	incrementQualityOnce := func(i Item) Item {i.Quality = i.Quality+1; return i}
	
	assertEquals("does nothing if item quality less than fifty", 
		     itemQualityLessThanFifty(Item{}),
		     itemQualityLessThanFifty(Item{}))
	assertEquals("does nothing if not backstage",
		     itemNameNotBackstage(itemQualityGreaterThanFifty(Item{})),
		     itemNameNotBackstage(itemQualityGreaterThanFifty(Item{})))
	assertEquals("does nothing if backstage but sell in too large",
		     itemLargeSellIn(itemNameNotBackstage(itemQualityGreaterThanFifty(Item{}))),
		     itemLargeSellIn(itemQualityGreaterThanFifty(Item{})))
	assertEquals("increments quality twice",
		     incrementQualityTwice(itemSellinLessThanSix(itemNameIsBackstage(itemQualityLessThanFifty(Item{})))),
		     itemSellinLessThanSix(itemNameIsBackstage(itemQualityLessThanFifty(Item{}))))
	assertEquals("increments quality once",
		     incrementQualityOnce(itemSellinLessThanSix(itemNameIsBackstage(itemQualityLessThanFifty(Item{})))),
		     itemSellinBetweenSixAndEleven(itemNameIsBackstage(itemQualityLessThanFifty(Item{}))))
		     
}

func main() {
	test("Before", Before)
	test("After", After)
	excercise()
}
