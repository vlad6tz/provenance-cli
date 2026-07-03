package tui

import (
	"provenance/internal/c2pa"
)

type errMsg error

type Model struct {
	FilePath string
	Report   *c2pa.ProvenanceReport
	Err      error
	Styles   UIStyles
	Parser   *c2pa.Parser
}

func NewModel(filepath string)
