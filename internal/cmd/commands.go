package cmd

import "github.com/spf13/cobra"

func SetCommands(c *cobra.Command) {

	c.AddCommand(getApplication())

	c.AddCommand(fetchAssignments())
	c.AddCommand(fetchApplications())
	c.AddCommand(fetchAssignmentsByApplication())

	c.AddCommand(createApplication())
	c.AddCommand(addAssignment())
	c.AddCommand(authenticate())
	c.AddCommand(register())

	c.AddCommand(modifyApplication())

	c.AddCommand(deleteApplication())
	c.AddCommand(removeAssignment())
}
