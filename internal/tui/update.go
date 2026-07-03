package tui

import (
	"provenance/internal/c2pa"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case *c2pa.ProvenanceReport:
		m.Report = msg
		return m, nil

	case errMsg:
		m.Err = msg
		return m, nil
	}

	return m, nil
}
