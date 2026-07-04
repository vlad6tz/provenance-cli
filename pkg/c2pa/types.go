// Package c2pa provides types and parsers for reading C2PA (Coalition for
// Content Provenance and Authenticity) provenance manifests embedded in
// digital image files.
//
// It reads manifest data such as creator information, timestamps, and
// cryptographic signatures to help inspect the origin of digital content.
//
// This is a read-only, unverified reader: it reports what the file claims
// about its provenance, but does not validate certificate chains against
// the C2PA trust list.
package c2pa

import "time"

// ValidationStatus represents the state of a C2PA signature or manifest.
type ValidationStatus string

const (
	StatusValid      ValidationStatus = "VALID"
	StatusInvalid    ValidationStatus = "INVALID"
	StatusUnverified ValidationStatus = "UNVERIFIED"
	StatusNone       ValidationStatus = "NONE"
)

// ProvenanceReport contains the parsed C2PA manifest data and file metadata.
type ProvenanceReport struct {
	HasManifest     bool
	SignatureStatus ValidationStatus
	SignatureIssuer string
	CreatorApp      string
	Timestamp       *time.Time
	FileName        string
	FileSize        int64
	MimeType        string
	SHA256          string
}
