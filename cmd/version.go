package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const AppVersion = "1.0.0" // hardcoded for now, and also why not  hihi

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "--v", "--version"},
	Short:   "Show the app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("prov%v", AppVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
