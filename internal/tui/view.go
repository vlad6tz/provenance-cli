package tui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	// Header
	b.WriteString("\n")
	b.WriteString(m.Styles.AppName.Render("Provenance"))
	b.WriteString(m.Styles.FileTitle.Render(m.FilePath) + "\n\n")

	// Loading
	if m.Report == nil && m.Err == nil {
		b.WriteString(" " + m.Styles.HelpText.Render("Reading file...") + "\n")
		return b.String()
	}

	// Error handling
	if m.Err != nil {
		b.WriteString("  " + m.Styles.StatusErr.Render("Error") + "  " + m.Err.Error() + "\n\n")
		b.WriteString(m.Styles.HelpText.Render("press q to exit"))
		return b.String()
	}

	// Main section
	b.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.MetaLabel.Render("size"), m.Styles.MetaValue.Render(fmt.Sprintf("%d bytes", m.Report.FileSize))))
	b.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.MetaLabel.Render("type"), m.Styles.MetaValue.Render(m.Report.MimeType)))
	b.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.MetaLabel.Render("sha256"), m.Styles.MetaHash.Render(m.Report.SHA256)))
	b.WriteString("\n")

	if !m.Report.HasManifest {
		b.WriteString("  " + m.Styles.StatusWarn.Render("No C2PA provenance found.") + "\n\n")

		b.WriteString(m.Styles.DisclaimerText.Render(
			m.Styles.DisclaimerText.Render("This does NOT necessarily mean the image is fake or AI-generated."),
		) + "\n")

		b.WriteString(m.Styles.HelpText.Render("press q to quit"))
		return b.String()
	}

	b.WriteString("  " + m.Styles.StatusValid.Render("Manifest Store Detected") + "\n")

	switch m.Report.SignatureStatus {
	case "VALID":
		b.WriteString("  " + m.Styles.StatusValid.Render("Cryptographic Signature Verified") + "\n")
	case "UNVERIFIED":
		b.WriteString("  " + m.Styles.StatusWarn.Render("Signature Unverified") + "\n")
	default:
		b.WriteString("  " + m.Styles.StatusErr.Render("Signature Corrupted / Invalid") + "\n")
	}
	b.WriteString("\n")

	b.WriteString(m.Styles.SectionHeader.Render("  METADATA SUMMARY") + "\n")

	if m.Report.CreatorApp != "" {
		b.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.MetaLabel.Render("application"), m.Styles.MetaValue.Render(m.Report.CreatorApp)))
	} else {
		b.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.MetaLabel.Render("application"), m.Styles.MetaHash.Render("unknown issuer")))
	}

	if m.Report.Timestamp != nil {
		formattedTime := m.Report.Timestamp.Format("2006-01-02 15:04 UTC")
		b.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.MetaLabel.Render("recorded at"), m.Styles.MetaValue.Render(formattedTime)))
	}

	b.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.MetaLabel.Render("assertions"), m.Styles.MetaValue.Render(fmt.Sprintf("%d signatures", m.Report.AssertionCount))))

	// Footer
	b.WriteString(m.Styles.HelpText.Render("  press q to quit") + "\n")

	return b.String()
}
