package tui

import (
	"github.com/vlad6tz/provenance-cli/pkg/c2pa"

	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type Model struct {
	FilePath string
	Report   *c2pa.ProvenanceReport
	Err      error
	Styles   UIStyles
	Parser   *c2pa.Parser
}

func NewModel(filepath string, parser *c2pa.Parser) Model {
	return Model{
		FilePath: filepath,
		Styles:   DefaultStyles(),
		Parser:   parser,
	}
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		report, err := m.Parser.ParseFile(m.FilePath)
		if err != nil {
			return errMsg(err)
		}
		return report
	}
}
