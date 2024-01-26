package main

import (
	"log"

	authCmd "github.com/charmingruby/clize/internal/auth/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func main() {
	authCmd.SetCommands(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
