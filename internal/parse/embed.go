package parse

import _ "embed"

// The members JSON files are embedded at compile time so that the scan
// path (which runs on end-user machines where parse/*.json are not on
// disk) can still answer "is this resource deprecated?" without a network
// round-trip. The files are produced by the weekly Resources workflow and
// committed to the repo; the embed happens against whatever is on disk at
// `go build` time, so binaries stay consistent with their source tree.
//
// The embedded bytes are raw JSON — callers unmarshal via their own type
// so this package doesn't leak the provider struct (unexported) out to
// other packages.

//go:embed aws-members.json
var awsMembersJSON []byte

//go:embed azurerm-members.json
var azureMembersJSON []byte

//go:embed google-members.json
var googleMembersJSON []byte

// EmbeddedMembers returns the raw members JSON for a provider by its
// short name (aws, azurerm/azure, google/gcp). Unknown names return nil;
// empty byte slices indicate an embedded-but-empty file, which in
// practice only happens if the members JSON was deleted before build.
func EmbeddedMembers(name string) []byte {
	switch name {
	case "aws":
		return awsMembersJSON
	case "azurerm", "azure":
		return azureMembersJSON
	case "google", "gcp":
		return googleMembersJSON
	default:
		return nil
	}
}
