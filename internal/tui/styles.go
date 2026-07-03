package tui

import "github.com/charmbracelet/lipgloss"

type UIStyles struct {
	Header        lipgloss.Style
	SubHeader     lipgloss.Style
	SuccessText   lipgloss.Style
	WarningText   lipgloss.Style
	ErrorText     lipgloss.Style
	MutedText     lipgloss.Style
	LabelText     lipgloss.Style
	ValueText     lipgloss.Style
	BorderLine    lipgloss.Style
	DisclaimerBox lipgloss.Style
}

func DefaultStyles() UIStyles {
	return UIStyles{
		Header:        lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Bold(true),  // Blue
		SubHeader:     lipgloss.NewStyle().Foreground(lipgloss.Color("246")),            // Gray
		SuccessText:   lipgloss.NewStyle().Foreground(lipgloss.Color("76")).Bold(true),  // Green
		WarningText:   lipgloss.NewStyle().Foreground(lipgloss.Color("220")),            // Yellow
		ErrorText:     lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true), // Red
		MutedText:     lipgloss.NewStyle().Foreground(lipgloss.Color("243")),            // Secondary Gray
		LabelText:     lipgloss.NewStyle().Foreground(lipgloss.Color("248")),            // Subtle white-gray
		ValueText:     lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true),
		BorderLine:    lipgloss.NewStyle().Foreground(lipgloss.Color("239")),
		DisclaimerBox: lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true, false, true, false).BorderForeground(lipgloss.Color("241")).Padding(0, 1),
	}
}
