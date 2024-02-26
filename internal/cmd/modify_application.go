package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
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
				terminal.PrintErrorResponse("application name is required")
				os.Exit(1)
			}

			if name == "" && context == "" {
				terminal.PrintErrorResponse("name or context is required")
				os.Exit(1)
			}

			if err := requests.ModifyApplication(appName, name, context); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&appName, "app name", "n", "", "app name")
	cmd.Flags().StringVarP(&name, "name", "m", "", "name")
	cmd.Flags().StringVarP(&context, "context", "c", "", "context")

	return cmd
}
