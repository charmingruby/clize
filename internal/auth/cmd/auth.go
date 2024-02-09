package cmd

import (
	"os"

	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/charmingruby/clize/pkg/requests"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func authenticate() *cobra.Command {
	var (
		username string
		password string
	)

	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticates an user",
		Run: func(cmd *cobra.Command, args []string) {
			if username == "" || password == "" {
				color.Red("username and password are required")
				os.Exit(1)
			}

			if err := requests.Auth(username, password, "/sign-in"); err != nil {
				cliui.PrintErrorResponse(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "password")

	return cmd
}
