package main

import (
	"log"

	"github.com/charmingruby/clize/internal/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func main() {
	cmd.SetCommands(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
