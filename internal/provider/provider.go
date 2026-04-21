// Package provider defines canonical provider names and alias normalisation.
// All provider-string comparisons in pike should go through Normalise so that
// alias handling (gcp↔google, azure↔azurerm) lives in exactly one place.
package provider

import "strings"

const (
	AWS    = "aws"
	Azure  = "azurerm"
	Google = "google"
	GCP    = "gcp"
)

// Normalise maps the various aliases pike accepts to the canonical short name.
// Returns an empty string for unrecognised input.
func Normalise(p string) string {
	switch strings.ToLower(p) {
	case AWS:
		return AWS
	case "azure", Azure:
		return Azure
	case GCP, Google:
		return Google
	default:
		return ""
	}
}
