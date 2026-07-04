package tests

import (
	"testing"

	"github.com/vlad6tz/provenance-cli/pkg/c2pa"
)

func TestVerifier_NoManifest(t *testing.T) {
	v := c2pa.NewVerifier()
	report := &c2pa.ProvenanceReport{HasManifest: false}
	err := v.VerifyManifestSgn(report)
	if err != c2pa.ErrNoManifest {
		t.Errorf("expected ErrNoManifest, got %v", err)
	}
}

func TestVerifier_InvalidSignature(t *testing.T) {
	v := c2pa.NewVerifier()
	report := &c2pa.ProvenanceReport{
		HasManifest:     true,
		SignatureStatus: c2pa.StatusInvalid,
	}
	err := v.VerifyManifestSgn(report)
	if err != c2pa.ErrValidationFailed {
		t.Errorf("expected ErrValidationFailed, got %v", err)
	}
}

func TestVerifier_Valid(t *testing.T) {
	v := c2pa.NewVerifier()
	report := &c2pa.ProvenanceReport{
		HasManifest:     true,
		SignatureStatus: c2pa.StatusValid,
	}
	err := v.VerifyManifestSgn(report)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestVerifier_Unverified(t *testing.T) {
	v := c2pa.NewVerifier()
	report := &c2pa.ProvenanceReport{
		HasManifest:     true,
		SignatureStatus: c2pa.StatusUnverified,
	}
	err := v.VerifyManifestSgn(report)
	if err != nil {
		t.Errorf("expected no error for unverified status, got %v", err)
	}
}
