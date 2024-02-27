package terminal

import "fmt"

func SymbolPattern() string {
	return "~#>"
}

func Padding() string {
	return fmt.Sprintf("%s   ", SymbolPattern())
}

func Header() {
	fmt.Print(Padding())
	boldWhite.Printf("C l i z e\n")
}

func Title(title string) {
	fmt.Print(Padding())
	boldWhite.Printf("[ %s ]\n", title)
}

func Gap() {
	println(SymbolPattern())
}

func Content(content, variant string) {
	print(Padding())
	switch variant {
	case "danger":
		boldRed.Printf("%s\n", content)
	case "ldanger":
		lightenRed.Printf("%s\n", content)
	case "success":
		boldGreen.Printf("%s\n", content)
	case "lsuccess":
		lightenGreen.Printf("%s\n", content)
	case "white":
		boldWhite.Printf("%s\n", content)
	default:
		fmt.Printf("%s\n", content)
	}
}

func ContentKeyValue(key, value string) {
	print(Padding())
	boldWhite.Printf("%s: ", key)
	fmt.Printf("%s\n", value)
}

func Footer() {
	println(Padding())
}
