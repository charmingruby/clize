package main

import (
	"log"

	appCmd "github.com/charmingruby/clize/internal/app/cmd"
	authCmd "github.com/charmingruby/clize/internal/auth/cmd"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func main() {
	authCmd.SetCommands(RootCmd)
	appCmd.SetCommands(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
