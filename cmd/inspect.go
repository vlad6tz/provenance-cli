package cmd

import (
	"fmt"
	"os"
	"provenance/internal/c2pa"
	"provenance/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var inseptCmd = &cobra.Command{
	Use:     "inpect [image path]",
	Aliases: []string{"i", "ins"},
	Short:   "Inspect C2PA provenance in a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parser := c2pa.NewParser()
		model := tui.NewModel(args[0], parser)

		p := tea.NewProgram(model, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "error rendering UI instance %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(inseptCmd)
}
