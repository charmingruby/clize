package terminal

import "github.com/fatih/color"

func PrintServerError() {
	print(LimitatorPattern())
	color.Red("Server is closed")
}

func PrintErrorResponse(err error) {
	Padding()
	color.Red(err.Error())
}
