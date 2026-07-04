# Provenance CLI

A command-line tool for inspecting and verifying C2PA (Coalition for Content Provenance and Authenticity) provenance metadata in digital files. It reads manifest data such as creator information, timestamps, and cryptographic signatures to help validate the origin and integrity of digital content.

## Installation

**Requirements:** Go 1.25 or later

```bash
git clone https://github.com/vlad6tz/provenance-cli
cd provenance-cli
go mod download
go build -o prov.exe .
```

## Usage

```
prov [command] [image path]
```

### Commands

| Command | Alias | Description |
|---|---|---|
| `inspect [path]` | `i`, `ins` | Parse C2PA metadata and open an interactive terminal UI |
| `verify [path]` | `v`, `ver` | Parse C2PA metadata and print a verification report |
| `version` | `vn` | Show the app version |

### Examples

```bash
prov inspect image.jpg
prov verify image.png
prov version
```

## Commands in Detail

### inspect

Parses the file's C2PA manifest and launches a Bubble Tea terminal UI. Displays file metadata (name, size, MIME type, SHA-256 hash), manifest status, signature status, signer identity, creator application, and timestamp.

### verify

Parses the file's C2PA manifest and prints a text-based report to stdout. Reports file name, whether a C2PA store was detected, signature status, and signer identity. Exits with an error on failure.

### Interpreting Results

**No manifest found** — The file has no C2PA provenance metadata. This does not mean the file is AI-generated or manipulated. Most real-world images lack C2PA data because the software or camera used does not support it, the metadata was stripped during re-encoding or upload, or the file predates the C2PA standard. Absence of provenance is not evidence of inauthenticity.

**Manifest found, signature unverified** — A C2PA manifest is present but the file contains no cryptographic signature, or the signer identity could not be read. This can happen with early C2PA implementations or when only minimal metadata was embedded.

**Signature verified** — A valid C2PA cryptographic signature was found and the signer identity was read. Note that the C2PA reader used by this tool reports what the file *claims* about its origin and full certificate chain validation against the C2PA trust list is not performed.

### version

Prints the current version string (`prov1.0.0`).

## Supported Formats

- JPEG / JPG
- PNG

Other image formats may be processed as JPEG as a fallback but are not guaranteed to work.


## Library Usage

This project can also be used as a Go library:

```go
import "github.com/vlad6tz/provenance-cli/pkg/c2pa"

parser := c2pa.NewParser()
report, err := parser.ParseFile("image.jpg")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Manifest: %v, Signed by: %s\n", report.HasManifest, report.SignatureIssuer)
```

## Testing

```bash
go test ./tests/ -v
```

## Dependencies

- [spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - terminal styling
- [richardwooding/c2pa](https://github.com/richardwooding/c2pa) - C2PA manifest reading

## License

MIT
