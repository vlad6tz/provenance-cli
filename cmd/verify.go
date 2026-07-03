package cmd

import (
	"fmt"
	"os"
	"provenance/internal/c2pa"
	"provenance/internal/tui"

	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:     "verify [image path]",
	Aliases: []string{"v", "ver"},
	Short:   "Verify C2PA provenance in a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parser := c2pa.NewParser()
		styles := tui.DefaultStyles()

		report, err := parser.ParseFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, styles.StatusErr.Render(fmt.Sprintf("Verification failure %v", err)))
			os.Exit(1)
		}

		if !report.HasManifest {
			fmt.Fprintln(os.Stdout, styles.StatusWarn.Render("No C2PA provenance found."))
			fmt.Fprintln(os.Stdout, styles.DisclaimerText.Render("This does NOT necessarily mean the image is fake or AI-generated."))
			return
		}

		fmt.Printf("File Name:        %s\n", report.FileName)
		fmt.Printf("C2PA Store:       %s\n", styles.StatusValid.Render("DETECTED"))
		fmt.Printf("Signature status: %s\n", report.SignatureStatus)
		if report.SignatureIssuer != "" {
			fmt.Printf("Certificate CA:   %s\n", report.SignatureIssuer)
		}
	},
}

func init() {
	RootCmd.AddCommand(verifyCmd)
}
