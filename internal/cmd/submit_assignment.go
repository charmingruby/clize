package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/spf13/cobra"
)

func submitAssignment() *cobra.Command {
	var (
		appName         string
		assignmentTitle string
	)

	// as-sub
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(AssignmentActor, terminal.SubmitCmd),
		Short: "Submits an assignment from an application",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" && assignmentTitle == "" {
				terminal.PrintErrorResponse("application name and assignment title are required")
				os.Exit(1)
			}

			if err := requests.SubmitAssignment(appName, assignmentTitle); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&appName, "app name", "n", "", "app name")
	cmd.Flags().StringVarP(&assignmentTitle, "assignment title", "t", "", "assignment title")

	return cmd
}
