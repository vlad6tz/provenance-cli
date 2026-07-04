package c2pa

import "errors"

// Sentinel errors returned by Verifier.
var (
	ErrValidationFailed = errors.New("cryptographic validation sequence signature breakdown")
	ErrNoManifest       = errors.New("cannot execute signature validation without an active manifest payload")
)

// Verifier performs basic consistency checks on a ProvenanceReport.
type Verifier struct{}

// NewVerifier creates a new Verifier.
func NewVerifier() *Verifier {
	return &Verifier{}
}

// VerifyManifestSgn checks the manifest and signature status of a report.
// It returns ErrNoManifest if no manifest is present, ErrValidationFailed
// if the signature is explicitly marked invalid, or nil otherwise.
func (v *Verifier) VerifyManifestSgn(report *ProvenanceReport) error {
	if !report.HasManifest {
		return ErrNoManifest
	}
	if report.SignatureStatus == StatusInvalid {
		return ErrValidationFailed
	}
	return nil
}
