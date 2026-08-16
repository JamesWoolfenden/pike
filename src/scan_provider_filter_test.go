package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

// GetPermissionBag's provider filter must accept the alias forms
// documented for users ("gcp", "azure") as well as the canonical HCL
// provider prefixes ("google", "azurerm"). Previously the filter compared
// the raw flag value directly against the resource's canonical provider
// prefix, so "gcp"/"azure" never matched anything and silently produced an
// empty result.
func TestGetPermissionBag_ProviderAliasFilter(t *testing.T) {
	t.Parallel()

	resources := []pike.ResourceV2{
		{
			TypeName:   "resource",
			Name:       "google_compute_instance",
			Provider:   "google",
			Attributes: []string{"name", "machine_type", "zone"},
		},
		{
			TypeName:   "resource",
			Name:       "azurerm_key_vault",
			Provider:   "azurerm",
			Attributes: []string{"name", "resource_group"},
		},
	}

	tests := []struct {
		name      string
		prov      string
		wantGCP   bool
		wantAZURE bool
	}{
		{name: "gcp alias", prov: "gcp", wantGCP: true, wantAZURE: false},
		{name: "google canonical", prov: "google", wantGCP: true, wantAZURE: false},
		{name: "azure alias", prov: "azure", wantGCP: false, wantAZURE: true},
		{name: "azurerm canonical", prov: "azurerm", wantGCP: false, wantAZURE: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bag := pike.GetPermissionBag(resources, tt.prov, true)

			if gotGCP := len(bag.GCP) > 0; gotGCP != tt.wantGCP {
				t.Errorf("prov=%q: GCP permissions present = %v, want %v (bag=%+v)", tt.prov, gotGCP, tt.wantGCP, bag)
			}
			if gotAZURE := len(bag.AZURE) > 0; gotAZURE != tt.wantAZURE {
				t.Errorf("prov=%q: AZURE permissions present = %v, want %v (bag=%+v)", tt.prov, gotAZURE, tt.wantAZURE, bag)
			}
		})
	}
}
