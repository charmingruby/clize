package cmd

import (
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/spf13/cobra"
)

func deleteApplication() *cobra.Command {
	var (
		name string
	)

	// appdel
	cmd := &cobra.Command{
		Use:   terminal.CommandWrapper(ApplicationActor, terminal.DelCmd),
		Short: "Deletes an application",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				terminal.PrintErrorResponse("name is required")
				os.Exit(1)
			}

			if err := requests.DeleteApplication(name); err != nil {
				terminal.PrintErrorResponse(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name")

	return cmd
}
