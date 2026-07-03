package c2pa

import "errors"

var (
	ErrValidationFailed = errors.New("cryptographic validation sequence signature breakdown")
	ErrNoManifest       = errors.New("cannot execute signature validation without an active manifest payload")
)

type Verifier struct{}

func NewVerifier() *Verifier {
	return &Verifier{}
}

func (v *Verifier) VerifyManifestSgn(report *ProvenanceReport) error {
	if !report.HasManifest {
		return ErrNoManifest
	}
	if report.SignatureStatus == StatusInvalid {
		return ErrValidationFailed
	}
	return nil
}
