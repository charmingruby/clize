package cmd

import (
	"fmt"
	"os"

	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/charmingruby/clize/pkg/cmd"
	"github.com/charmingruby/clize/pkg/requests"
	"github.com/spf13/cobra"
)

func fetchApplications() *cobra.Command {

	// appftc
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s%s", ApplicationActor, cmd.FetchCmd),
		Short: "Fetch all applications",
		Run: func(cmd *cobra.Command, args []string) {
			if err := requests.FetchApplications(); err != nil {
				cliui.PrintErrorResponse(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
