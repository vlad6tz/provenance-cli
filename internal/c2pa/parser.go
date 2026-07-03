package c2pa

import (
	"context"
	"fmt"
	"os"
	"provenance/internal/utils"
	"strings"

	"github.com/richardwooding/c2pa"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFile(path string) (*ProvenanceReport, error) {
	meta, err := utils.GetFileMetadata(path)
	if err != nil {
		return nil, fmt.Errorf("failed gathering file meta: %w", err)
	}
	hash, err := utils.ComputeSHA256(path)
	if err != nil {
		return nil, fmt.Errorf("failed calculating file hash: %w", err)
	}

	report := &ProvenanceReport{
		FileName: meta.Name,
		FileSize: meta.Size,
		MimeType: meta.MimeType,
		SHA256:   hash,
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed opening file for manifest parsing: %w", err)
	}
	defer file.Close()

	// Resolve the correct c2pa.Container constant using the file's MIME type
	var container c2pa.Container
	switch {
	case strings.Contains(meta.MimeType, "jpeg") || strings.Contains(meta.MimeType, "jpg"):
		container = c2pa.JPEG
	case strings.Contains(meta.MimeType, "png"):
		container = c2pa.PNG
	default:
		// Fall back to JPEG processing if it's an unrecognized MIME wrapper type
		container = c2pa.JPEG
	}

	info := c2pa.Read(context.Background(), container, file)

	infoStr := fmt.Sprintf("%v", info)

	if infoStr == "{}" || infoStr == "" {
		report.HasManifest = false
		report.SignatureStatus = StatusNone
		return report, nil
	}

	report.HasManifest = true
	report.SignatureStatus = StatusUnverified
	report.AssertionCount = 0

	return report, nil
}
