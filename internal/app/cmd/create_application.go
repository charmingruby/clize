package cmd

import (
	"fmt"
	"os"

	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/charmingruby/clize/pkg/cmd"
	"github.com/charmingruby/clize/pkg/requests"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func createApplication() *cobra.Command {
	var (
		name    string
		context string
	)

	// appcrt
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s%s", ApplicationActor, cmd.CreateCmd),
		Short: "Creates a new application",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" && context == "" {
				color.Red("name and context are required")
				os.Exit(1)
			}

			if err := requests.CreateApplication(name, context); err != nil {
				cliui.PrintErrorResponse(err)
				os.Exit(1)
			}

		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name")
	cmd.Flags().StringVarP(&context, "context", "c", "", "context")

	return cmd
}
