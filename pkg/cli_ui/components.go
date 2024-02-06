package cliui

import "fmt"

func Header() string {
	return fmt.Sprintf("%s C l i z e", LimitatorPattern())
}

func Title(title string) string {
	return fmt.Sprintf("%s( %s )", Padding(), title)
}

func Padding() string {
	return fmt.Sprintf("%s  ", SymbolPattern())
}

func Gap() string {
	return fmt.Sprint(SymbolPattern())
}

func Content(content string) string {
	return fmt.Sprintf("%s  %s", SymbolPattern(), content)
}

func Footer() string {
	return LimitatorPattern()
}
