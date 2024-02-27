package terminal

import "fmt"

func PrintServerError() {
	Content(
		"Server is closed",
		"danger",
	)
}

func PrintErrorResponse(msg string) {
	Header()
	Gap()

	Content(
		fmt.Sprintf("❌ %s", msg),
		"danger",
	)

	Gap()
	Footer()
}

func PrintSuccessMsgResponse(msg string) {
	Header()
	Gap()

	Content(
		fmt.Sprintf("✅ %s", msg),
		"success",
	)

	Gap()
	Footer()
}
