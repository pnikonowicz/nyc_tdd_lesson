package main

func printBanner() string {
	return "BANNER\n"
}

func Before(amount int, name string) string {
	banner := printBanner()

	//print details
	banner += "name:" + name + "\n"
	banner += "amount:" + string(amount) + "\n"

	return banner
}

func After(amount int, name string) string {
	printDetails := func(banner string) string {
		banner += "name:" + name + "\n"
		banner += "amount:" + string(amount) + "\n"
		return banner
	}

	banner := printBanner()
	details := printDetails(banner)

	return banner + details
}

func Excercise(elements []string, name string, amount string) string {
	owingStr := ""
	print := func(str string) {
		owingStr += str + "\n"
	}

	//banner
	print("****Banner")

	//calculate outstanding
	for _, element := range elements {
		print("Element: " + element)
	}

	//print details
	print("name:" + name)
	print("amount:" + amount)

	return owingStr
}
