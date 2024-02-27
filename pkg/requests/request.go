package requests

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/charmingruby/clize/pkg/terminal"
	jwt "github.com/charmingruby/clize/pkg/token"
)

func doRequest(
	method, path string, body io.Reader, auth bool,
) (*http.Response, error) {
	// Comment this block to install
	// if err := godotenv.Load(); err != nil {
	// 	terminal.PrintErrorResponse(".env don't exists")
	// 	os.Exit(1)
	// }

	// serverUrl, ok := os.LookupEnv("SERVER_URL")

	// if !ok {
	// 	terminal.PrintErrorResponse("SERVER_URL not found on .env")
	// 	os.Exit(1)
	// }
	//

	// Change the server URL to the current server url and run make install
	url := fmt.Sprintf("%s%s", "https://us1-ruling-gelding-37317.upstash.io", path)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if auth {
		token, err := readCacheToken()

		if err != nil {
			log.Println("Cannot read cache token.")
			return nil, err
		}

		if isTokenValid := jwt.NewJwtService().ValidateToken(token); !isTokenValid {
			terminal.PrintErrorResponse("cannot validate token")
			os.Exit(1)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	return http.DefaultClient.Do(req)
}
