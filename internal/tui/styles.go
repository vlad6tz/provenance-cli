package tui

import "github.com/charmbracelet/lipgloss"

type UIStyles struct {
	AppName        lipgloss.Style
	FileTitle      lipgloss.Style
	MetaLabel      lipgloss.Style
	MetaValue      lipgloss.Style
	MetaHash       lipgloss.Style
	StatusValid    lipgloss.Style
	StatusWarn     lipgloss.Style
	StatusErr      lipgloss.Style
	SectionHeader  lipgloss.Style
	DisclaimerText lipgloss.Style
	HelpText       lipgloss.Style
}

func DefaultStyles() UIStyles {
	return UIStyles{
		AppName:   lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).MarginRight(2),
		FileTitle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true),
		MetaLabel: lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Width(14),
		MetaValue: lipgloss.NewStyle().Foreground(lipgloss.Color("252")),
		MetaHash:  lipgloss.NewStyle().Foreground(lipgloss.Color("239")),

		StatusValid: lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Bold(true),
		StatusWarn:  lipgloss.NewStyle().Foreground(lipgloss.Color("216")),
		StatusErr:   lipgloss.NewStyle().Foreground(lipgloss.Color("203")).Bold(true),

		SectionHeader: lipgloss.NewStyle().Foreground(lipgloss.Color("111")).Bold(true).MarginTop(1).MarginBottom(1),

		DisclaimerText: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color("237")).
			PaddingLeft(2).
			Foreground(lipgloss.Color("241")),

		HelpText: lipgloss.NewStyle().Foreground(lipgloss.Color("237")).MarginTop(2),
	}
}
