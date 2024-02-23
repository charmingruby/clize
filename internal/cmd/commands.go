package cmd

import "github.com/spf13/cobra"

func SetCommands(c *cobra.Command) {
	c.AddCommand(ping())

	c.AddCommand(getApplication())
	c.AddCommand(fetchAssignmentsByApplication())

	c.AddCommand(fetchAssignments())
	c.AddCommand(fetchApplications())

	c.AddCommand(createApplication())
	c.AddCommand(addAssignment())
	c.AddCommand(authenticate())
	c.AddCommand(register())

	c.AddCommand(modifyApplication())
	c.AddCommand(modifyAssignment())
	c.AddCommand(submitAssignment())

	c.AddCommand(deleteApplication())
	c.AddCommand(removeAssignment())
}
