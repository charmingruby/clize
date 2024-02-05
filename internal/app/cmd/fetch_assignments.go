package cmd

import (
	"fmt"
	"os"

	"github.com/charmingruby/clize/pkg/requests"
	"github.com/spf13/cobra"
)

func fetchAssignments() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "assignments",
		Short: "Fetch all assignments",
		Run: func(cmd *cobra.Command, args []string) {
			if err := requests.FetchAssignments(); err != nil {
				fmt.Printf("%x", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
