package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func removeAssignment() *cobra.Command {
	var (
		appName        string
		assignmentName string
	)

	// as-rm
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(AssignmentActor, terminal.RemoveCmd),
		Short: "Removes an assignment from an application",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" && assignmentName == "" {
				color.Red("app and assignment names are required")
				os.Exit(1)
			}

			if err := requests.RemoveAssignment(appName, assignmentName); err != nil {
				terminal.PrintErrorResponse(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&appName, "app name", "n", "", "app name")
	cmd.Flags().StringVarP(&assignmentName, "assignment name", "a", "", "assignment name")

	return cmd
}
