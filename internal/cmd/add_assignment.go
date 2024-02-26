package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/spf13/cobra"
)

func addAssignment() *cobra.Command {
	var (
		appName     string
		title       string
		description string
	)

	// as-add
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(AssignmentActor, terminal.AddCmd),
		Short: "Creates a new assignment",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" {
				terminal.PrintErrorResponse("application name is required")
				os.Exit(1)
			}

			if title == "" && description == "" {
				terminal.PrintErrorResponse("name and description are required")
				os.Exit(1)
			}

			if err := requests.AddAssignment(appName, title, description); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}

		},
	}

	cmd.Flags().StringVarP(&appName, "appName", "n", "", "appName")
	cmd.Flags().StringVarP(&title, "title", "t", "", "title")
	cmd.Flags().StringVarP(&description, "description", "d", "", "description")

	return cmd
}
