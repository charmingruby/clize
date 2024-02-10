package terminal

import "fmt"

func SymbolPattern() string {
	return "~"
}

func LimitatorPattern() string {
	return fmt.Sprintf("%s# ", SymbolPattern())
}

func paddingPattern() string {
	return fmt.Sprintf("%s    ", SymbolPattern())
}

func Header() {
	fmt.Print(LimitatorPattern())
	BoldWhite.Printf(" C l i z e\n")
}

func Title(title string) {
	BoldWhite.Printf("%s( %s )\n", paddingPattern(), title)
}

func Tab(tabs int) {
	tab := "  "

	var tabsToIdent string

	for i := 0; i < tabs; i++ {
		tabsToIdent = tabsToIdent + tab
	}

	print(tabsToIdent)
}

func Padding() {
	print(paddingPattern())
}

func Gap() {
	println(SymbolPattern())
}

func Content(content string) {
	fmt.Printf("%s%s\n", paddingPattern(), content)
}

func Footer() {
	println(LimitatorPattern())
}
