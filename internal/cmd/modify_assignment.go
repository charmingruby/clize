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
		appName         string
		assignmentTitle string
		title           string
		description     string
	)

	// as-mod
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(AssignmentActor, terminal.ModifyCmd),
		Short: "Modifies an assignment",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" && assignmentTitle == "" {
				color.Red("app name and assignment title are required")
				os.Exit(1)
			}

			if title == "" && description == "" {
				color.Red("title and description are required")
				os.Exit(1)
			}

			if err := requests.ModifyAssignment(appName, assignmentTitle, title, description); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&appName, "app name", "n", "", "app name")
	cmd.Flags().StringVarP(&assignmentTitle, "assignment title", "t", "", "assignment title")
	cmd.Flags().StringVarP(&title, "title", "x", "", "title")
	cmd.Flags().StringVarP(&description, "description", "y", "", "description")

	return cmd
}
