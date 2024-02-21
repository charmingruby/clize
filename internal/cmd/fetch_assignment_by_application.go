package cmd

import (
	"fmt"
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func fetchAssignmentsByApplication() *cobra.Command {
	var (
		name string
	)

	actors := fmt.Sprintf("%s-%s", ApplicationActor, AssignmentActor)

	// app-as-get
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(actors, terminal.GetCmd),
		Short: "Get all assignments of an application",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				color.Red("name are required")
				os.Exit(1)
			}

			if err := requests.FetchAssignmentsByApplication(name); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name")

	return cmd
}
