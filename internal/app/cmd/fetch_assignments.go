package cmd

import (
	"os"

	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/charmingruby/clize/pkg/requests"
	"github.com/spf13/cobra"
)

func fetchAssignments() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "aas",
		Short: "Fetch all assignments",
		Run: func(cmd *cobra.Command, args []string) {
			if err := requests.FetchAssignments(); err != nil {
				cliui.PrintErrorResponse(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
