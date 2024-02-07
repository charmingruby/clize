package cliui

import "github.com/fatih/color"

func PrintServerError() {
	print(LimitatorPattern())
	color.Red("Server is closed")
}

func PrintErrorResponse(err error) {
	print(Padding())
	color.Red(err.Error())
}
