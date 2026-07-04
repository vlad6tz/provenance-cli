package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vlad6tz/provenance-cli/internal/utils"
)

func TestGetFileMetadata_NonExistent(t *testing.T) {
	_, err := utils.GetFileMetadata(filepath.Join(os.TempDir(), "nonexistent-file-12345"))
	if err != utils.ErrFileNotFound {
		t.Errorf("expected ErrFileNotFound, got %v", err)
	}
}

func TestGetFileMetadata_EmptyFile(t *testing.T) {
	f, err := os.CreateTemp("", "empty-test-*.tmp")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.Close()

	_, err = utils.GetFileMetadata(f.Name())
	if err != utils.ErrEmptyFile {
		t.Errorf("expected ErrEmptyFile, got %v", err)
	}
}

func TestGetFileMetadata_Valid(t *testing.T) {
	f, err := os.CreateTemp("", "test-*.tmp")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.WriteString("hello world")
	f.Close()

	meta, err := utils.GetFileMetadata(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if meta.Size != 11 {
		t.Errorf("expected size 11, got %d", meta.Size)
	}
	if meta.Name != filepath.Base(f.Name()) {
		t.Errorf("expected name %s, got %s", filepath.Base(f.Name()), meta.Name)
	}
}

func TestComputeSHA256(t *testing.T) {
	f, err := os.CreateTemp("", "hash-test-*.tmp")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.WriteString("hello world")
	f.Close()

	hash, err := utils.ComputeSHA256(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	if hash != expected {
		t.Errorf("expected %s, got %s", expected, hash)
	}
}

func TestComputeSHA256_NonExistent(t *testing.T) {
	_, err := utils.ComputeSHA256(filepath.Join(os.TempDir(), "nonexistent-hash-file"))
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}
