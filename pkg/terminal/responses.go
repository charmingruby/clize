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

func PrintNotFoundResponse(identifier string) {
	Header()
	Gap()

	Padding()
	BoldRed.Printf("%s not found\n", identifier)

	Gap()
	Footer()
}

func PrintSuccessMsgResponse(msg string) {
	Header()
	Gap()

	Padding()
	BoldGreen.Printf("✅ %s\n", msg)

	Gap()
	Footer()
}
