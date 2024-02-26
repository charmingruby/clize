package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/spf13/cobra"
)

func createApplication() *cobra.Command {
	var (
		name    string
		context string
	)

	// appcrt
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(ApplicationActor, terminal.CreateCmd),
		Short: "Creates a new application",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" && context == "" {
				terminal.PrintErrorResponse("name and context are required")
				os.Exit(1)
			}

			if err := requests.CreateApplication(name, context); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}

		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name")
	cmd.Flags().StringVarP(&context, "context", "c", "", "context")

	return cmd
}
