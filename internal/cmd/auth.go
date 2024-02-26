package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
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
				terminal.PrintErrorResponse("username and password are required")
				os.Exit(1)
			}

			if err := requests.Auth(username, password); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "password")

	return cmd
}
