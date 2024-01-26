package cmd

import (
	"log"
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/spf13/cobra"
)

func authenticate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticates an user with Auth0",
		Run: func(cmd *cobra.Command, args []string) {
			if err := requests.Auth("/login"); err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
