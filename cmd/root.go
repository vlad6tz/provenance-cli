package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "provenance",
	Short: "Inspect C2PA provenance data in digital files",
	Long: `Provenance is a CLI tool for inspecting and verifying C2PA (Coalition for Content Provenance and Authenticity) metadata embedded in digital files.

It reads manifest data such as creator info, timestamps, and cryptographic signatures to help validate the origin and integrity of digital content.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime exception: %v\n", err)
		os.Exit(1)
	}
}
