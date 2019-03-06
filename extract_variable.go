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
