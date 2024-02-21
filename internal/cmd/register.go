package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func register() *cobra.Command {
	var (
		username string
		email    string
		password string
	)

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Registers a new user",
		Run: func(cmd *cobra.Command, args []string) {
			if username == "" && email == "" && password == "" {
				color.Red("username, email and password are required")
				os.Exit(1)
			}

			if err := requests.Register(username, email, password); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "username")
	cmd.Flags().StringVarP(&email, "email", "e", "", "email")
	cmd.Flags().StringVarP(&password, "password", "p", "", "password")

	return cmd
}
