package cmd

import "github.com/spf13/cobra"

func SetCommands(c *cobra.Command) {
	c.AddCommand(createApplication())
	c.AddCommand(fetchAssignments())
	c.AddCommand(fetchApplications())
	c.AddCommand(getApplication())
	c.AddCommand(deleteApplication())
	c.AddCommand(modifyApplication())
	c.AddCommand(addAssignment())
	c.AddCommand(authenticate())
	c.AddCommand(register())
}
