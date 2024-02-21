package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func modifyAssignment() *cobra.Command {
	var (
		appName      string
		assignmentID string
		title        string
		description  string
	)

	// as-mod
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(AssignmentActor, terminal.ModifyCmd),
		Short: "Modifies an assignment",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" && assignmentID == "" {
				color.Red("app and assignment are required")
				os.Exit(1)
			}

			if title == "" && description == "" {
				color.Red("title and description are required")
				os.Exit(1)
			}

			if err := requests.ModifyAssignment(appName, assignmentID, title, description); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&appName, "app name", "a", "", "app name")
	cmd.Flags().StringVarP(&assignmentID, "assignment id", "i", "", "assignment id")
	cmd.Flags().StringVarP(&title, "title", "t", "", "title")
	cmd.Flags().StringVarP(&description, "description", "d", "", "description")

	return cmd
}
