package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func submitAssignment() *cobra.Command {
	var (
		appName      string
		assignmentID string
	)

	// as-sub
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(AssignmentActor, terminal.SubmitCmd),
		Short: "Submits an assignment from an application",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" && assignmentID == "" {
				color.Red("app name and assignment id are required")
				os.Exit(1)
			}

			if err := requests.SubmitAssignment(appName, assignmentID); err != nil {
				terminal.PrintErrorResponse(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&appName, "app name", "n", "", "app name")
	cmd.Flags().StringVarP(&assignmentID, "assignment id", "i", "", "assignment id")

	return cmd
}
