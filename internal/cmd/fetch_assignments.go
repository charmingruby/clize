package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/spf13/cobra"
)

func fetchAssignments() *cobra.Command {

	// asftc
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(AssignmentActor, terminal.FetchCmd),
		Short: "Fetch all assignments",
		Run: func(cmd *cobra.Command, args []string) {
			if err := requests.FetchAssignments(); err != nil {
				terminal.PrintErrorResponse(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
