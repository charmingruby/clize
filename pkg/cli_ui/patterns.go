package cliui

import "fmt"

func SymbolPattern() string {
	return "~"
}

func LimitatorPattern() string {
	return fmt.Sprintf("%s# ", SymbolPattern())
}
