package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/spf13/cobra"
)

func ping() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hc",
		Short: "Checks the health of the server",
		Run: func(cmd *cobra.Command, args []string) {
			//cmd
			if err := requests.Ping(); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	return cmd
}
