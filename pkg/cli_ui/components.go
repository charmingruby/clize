package cliui

import "fmt"

func Header() {
	fmt.Print(LimitatorPattern())
	BoldWhite.Printf(" C l i z e\n")
}

func Title(title string) {
	BoldWhite.Printf("%s( %s )\n", Padding(), title)
}

func Padding() string {
	return fmt.Sprintf("%s    ", SymbolPattern())
}

func Gap() string {
	return fmt.Sprint(SymbolPattern())
}

func Content(content string) string {
	return fmt.Sprintf("%s%s", Padding(), content)
}

func Footer() string {
	return LimitatorPattern()
}
