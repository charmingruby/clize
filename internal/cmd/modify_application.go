package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func modifyApplication() *cobra.Command {
	var (
		appName string
		name    string
		context string
	)

	// appmod
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(ApplicationActor, terminal.ModifyCmd),
		Short: "Modifies an application",
		Run: func(cmd *cobra.Command, args []string) {
			if appName == "" {
				color.Red("app are required")
				os.Exit(1)
			}

			if name == "" && context == "" {
				color.Red("name and context are required")
				os.Exit(1)
			}

			if err := requests.ModifyApplication(appName, name, context); err != nil {
				terminal.PrintErrorResponse(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&appName, "app name", "a", "", "app name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "name")
	cmd.Flags().StringVarP(&context, "context", "c", "", "context")

	return cmd
}
