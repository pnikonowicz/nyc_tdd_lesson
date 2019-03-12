package nyc_tdd_lesson

import (
	. "strings"
)

//For Page 124 Introduce Explaining Variable
func Before(platform, browser string, resize int, wasInitialized func() bool) {
	if Compare(ToUpper(platform), "MAC") > -1 &&
		Compare(ToUpper(browser), "IE") > -1 &&
		wasInitialized() && resize > 0 {
		// do something
	}
}

func After(platform, browser string, resize int, wasInitialized func() bool) {
	isMacOs := Compare(ToUpper(platform), "MAC") > -1
	isIEBrowser := Compare(ToUpper(browser), "IE") > -1
	wasResized := resize > 0

	if isMacOs &&
		isIEBrowser &&
		wasInitialized() &&
		wasResized {
		// do something
	}
}

func assert(expected bool, testFunction func() bool) {
	actual := testFunction()
	
	if expected != actual {
		panic(fmt.Sprintf("%q != %q", expected, actual))
	}
}

func main() {
	isMacOs := "MAC"
	isIeBrowser := "IE"
	wasResized := 1 // number > 0
	isInitialized := func() bool {return true}
	
	notMacOs := "not mac"
	notIeBrowser := "not ie"
	notResized := -1 // number <= 0
	notInitialized := func() bool {return false}
	
	assert(true, func() bool { Before(isMacOS, isIeBrowser, wasResized, isInitialized) })
	assert(false, func() bool { Before(notMacOS, isIeBrowser, wasResized, isInitialized) })
	assert(false, func() bool { Before(isMacOS, notIeBrowser, wasResized, isInitialized) })
	assert(false, func() bool { Before(isMacOS, isIeBrowser, notResized, isInitialized) })
	assert(false, func() bool { Before(isMacOS, isIeBrowser, wasResized, notInitialized) })
}
