package main

import (
	"log"

	cmd "github.com/tailwarden/iam-wiz/cmd"
)

func main() {
	if err := cmd.CliCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
