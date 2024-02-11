package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
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
		Short: "Creates a new application",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" {
				color.Red("app are required")
				os.Exit(1)
			}

			if title == "" && description == "" {
				color.Red("name and description are required")
				os.Exit(1)
			}

			if err := requests.AddAssignment(appName, title, description); err != nil {
				terminal.PrintErrorResponse(err)
				os.Exit(1)
			}

		},
	}

	cmd.Flags().StringVarP(&appName, "appName", "a", "", "appName")
	cmd.Flags().StringVarP(&title, "title", "t", "", "title")
	cmd.Flags().StringVarP(&description, "description", "d", "", "description")

	return cmd
}
