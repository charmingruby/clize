package cmd

import "github.com/spf13/cobra"

func SetCommands(c *cobra.Command) {
	c.AddCommand(authenticate())
	c.AddCommand(register())
}
