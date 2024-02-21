package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func getApplication() *cobra.Command {
	var (
		name string
	)

	// appget
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(ApplicationActor, terminal.GetCmd),
		Short: "Get an applications",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				color.Red("name are required")
				os.Exit(1)
			}

			if err := requests.GetApplication(name); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name")

	return cmd
}
