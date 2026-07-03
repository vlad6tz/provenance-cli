package c2pa

import "time"

type ValidationStatus string

const (
	StatusValid      ValidationStatus = "VALID"
	StatusInvalid    ValidationStatus = "INVALID"
	StatusUnverified ValidationStatus = "UNVERIFIED"
	StatusNone       ValidationStatus = "NONE"
)

type ProvenanceReport struct {
	HasManifest     bool
	SignatureStatus ValidationStatus
	SignatureIssuer string
	CreatorApp      string
	Timestamp       *time.Time
	AssertionCount  int
	FileName        string
	FileSize        int64
	MimeType        string
	SHA256          string
}
