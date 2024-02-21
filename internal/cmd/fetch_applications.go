package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/spf13/cobra"
)

func fetchApplications() *cobra.Command {

	// app-ftc
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(ApplicationActor, terminal.FetchCmd),
		Short: "Fetch all applications",
		Run: func(cmd *cobra.Command, args []string) {
			if err := requests.FetchApplications(); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	return cmd
}
