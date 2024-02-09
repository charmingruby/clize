package cmd

import (
	"fmt"
	"os"

	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/charmingruby/clize/pkg/cmd"
	"github.com/charmingruby/clize/pkg/requests"
	"github.com/spf13/cobra"
)

func fetchAssignments() *cobra.Command {

	// asftc
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s%s", AssignmentActor, cmd.FetchCmd),
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
