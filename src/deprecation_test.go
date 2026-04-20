package pike

import (
	"encoding/json"
	"strings"
	"testing"
)

// Test_DeprecationMessage hits the public lookup with a seeded cache so
// we don't depend on whatever the latest Resources workflow happened to
// embed into the binary.
func Test_DeprecationMessage(t *testing.T) {
	t.Cleanup(clearDeprecationsForTest)

	seedDeprecationForTest("aws", deprecationEntry{
		DeprecatedResources: map[string]string{
			"aws_retired_thing": "Deprecated: use aws_new_thing instead",
			"aws_silent_death":  "",
		},
		DeprecatedData: map[string]string{
			"aws_old_data": "Deprecated: query aws_new_data",
		},
	})

	tests := []struct {
		name     string
		provider string
		resource string
		isData   bool
		wantMsg  string
		wantOk   bool
	}{
		{
			name:     "resource hit with message",
			provider: "aws",
			resource: "aws_retired_thing",
			wantMsg:  "Deprecated: use aws_new_thing instead",
			wantOk:   true,
		},
		{
			name:     "resource hit without message",
			provider: "aws",
			resource: "aws_silent_death",
			wantMsg:  "",
			wantOk:   true,
		},
		{
			name:     "data hit",
			provider: "aws",
			resource: "aws_old_data",
			isData:   true,
			wantMsg:  "Deprecated: query aws_new_data",
			wantOk:   true,
		},
		{
			name:     "wrong kind: resource name queried as data",
			provider: "aws",
			resource: "aws_retired_thing",
			isData:   true,
			wantOk:   false,
		},
		{
			name:     "miss",
			provider: "aws",
			resource: "aws_live_and_well",
			wantOk:   false,
		},
		{
			name:     "unknown provider is safe",
			provider: "nonsense",
			resource: "aws_retired_thing",
			wantOk:   false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			msg, ok := DeprecationMessage(tc.provider, tc.resource, tc.isData)
			if ok != tc.wantOk {
				t.Errorf("ok = %v, want %v", ok, tc.wantOk)
			}
			if msg != tc.wantMsg {
				t.Errorf("msg = %q, want %q", msg, tc.wantMsg)
			}
		})
	}
}

// Test_DeprecationMessage_ProviderAliases exercises the gcp/google and
// azure/azurerm normalisation. Users invoke pike with either alias and
// the ResourceV2.Provider passed to warnIfDeprecated reflects that;
// the lookup must paper over the difference.
func Test_DeprecationMessage_ProviderAliases(t *testing.T) {
	t.Cleanup(clearDeprecationsForTest)

	seedDeprecationForTest("google", deprecationEntry{
		DeprecatedResources: map[string]string{"google_thing": "dead"},
	})
	seedDeprecationForTest("azurerm", deprecationEntry{
		DeprecatedResources: map[string]string{"azurerm_thing": "dead"},
	})

	for _, alias := range []string{"gcp", "google"} {
		msg, ok := DeprecationMessage(alias, "google_thing", false)
		if !ok || msg != "dead" {
			t.Errorf("alias %q: got (%q, %v), want (%q, %v)", alias, msg, ok, "dead", true)
		}
	}
	for _, alias := range []string{"azure", "azurerm"} {
		msg, ok := DeprecationMessage(alias, "azurerm_thing", false)
		if !ok || msg != "dead" {
			t.Errorf("alias %q: got (%q, %v), want (%q, %v)", alias, msg, ok, "dead", true)
		}
	}
}

// Test_warnIfDeprecated_Dedup confirms that the same deprecated resource
// declared multiple times in one scan only produces a single warned[]
// entry (and therefore only one log.Warn). This is the behaviour users
// actually see: noisy repeated warnings would be a regression.
func Test_warnIfDeprecated_Dedup(t *testing.T) {
	t.Cleanup(clearDeprecationsForTest)

	seedDeprecationForTest("aws", deprecationEntry{
		DeprecatedResources: map[string]string{
			"aws_retired_thing": "use aws_new_thing",
		},
	})

	warned := map[string]struct{}{}
	r := ResourceV2{
		Provider: "aws",
		Name:     "aws_retired_thing",
		TypeName: "resource",
	}

	warnIfDeprecated(r, warned)
	warnIfDeprecated(r, warned)
	warnIfDeprecated(r, warned)

	if got, want := len(warned), 1; got != want {
		t.Errorf("len(warned) = %d, want %d (dedup failed)", got, want)
	}
}

// Test_warnIfDeprecated_DistinguishesKind confirms that resource-and-data
// sharing a name (e.g. `google_storage_bucket`) each get their own warn
// bucket — if only the resource is deprecated, the data lookup stays
// silent, and vice versa.
func Test_warnIfDeprecated_DistinguishesKind(t *testing.T) {
	t.Cleanup(clearDeprecationsForTest)

	seedDeprecationForTest("google", deprecationEntry{
		DeprecatedResources: map[string]string{"google_bucket": "dead resource"},
		// Data form with same name is NOT deprecated.
	})

	warned := map[string]struct{}{}

	warnIfDeprecated(ResourceV2{
		Provider: "google", Name: "google_bucket", TypeName: "resource",
	}, warned)

	warnIfDeprecated(ResourceV2{
		Provider: "google", Name: "google_bucket", TypeName: "data",
	}, warned)

	if got, want := len(warned), 1; got != want {
		t.Errorf("len(warned) = %d, want %d (only resource kind is deprecated)", got, want)
	}
}

// Test_warnIfDeprecated_SkipsEmpty guards against accidental lookups on
// malformed or module-level resources whose Name or Provider hasn't been
// filled in. These must no-op cleanly rather than hitting the map with
// an empty key.
func Test_warnIfDeprecated_SkipsEmpty(t *testing.T) {
	warned := map[string]struct{}{}

	warnIfDeprecated(ResourceV2{Provider: "aws"}, warned)
	warnIfDeprecated(ResourceV2{Name: "aws_thing"}, warned)
	warnIfDeprecated(ResourceV2{}, warned)

	if len(warned) != 0 {
		t.Errorf("len(warned) = %d, want 0 (empty inputs must not register)", len(warned))
	}
}

// Test_Deprecated_SingleProvider covers the CLI-filtered path: asking
// for one provider returns exactly that provider's entries and nothing
// else. Unknown provider names return nil (not an error) because the
// CLI is informational.
func Test_Deprecated_SingleProvider(t *testing.T) {
	t.Cleanup(clearDeprecationsForTest)

	seedDeprecationForTest("aws", deprecationEntry{
		DeprecatedResources: map[string]string{"aws_old": "msg"},
	})
	seedDeprecationForTest("google", deprecationEntry{
		DeprecatedResources: map[string]string{"google_old": "msg"},
	})

	got := Deprecated("aws")
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if got[0].Provider != "aws" {
		t.Errorf("provider = %q, want aws", got[0].Provider)
	}

	if Deprecated("nonsense") != nil {
		t.Error("unknown provider should return nil")
	}
}

// Test_Deprecated_AllProviders covers the default CLI path: no provider
// filter walks all supported ones. Providers with nothing deprecated
// must be omitted entirely so the output doesn't show empty headers.
func Test_Deprecated_AllProviders(t *testing.T) {
	t.Cleanup(clearDeprecationsForTest)

	// aws has entries, azurerm is empty, google has a data-source-only entry.
	seedDeprecationForTest("aws", deprecationEntry{
		DeprecatedResources: map[string]string{"aws_old": "msg"},
	})
	seedDeprecationForTest("azurerm", deprecationEntry{})
	seedDeprecationForTest("google", deprecationEntry{
		DeprecatedData: map[string]string{"google_old_data": "msg"},
	})

	got := Deprecated("")
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2 (azurerm should be omitted as empty)", len(got))
	}

	providers := []string{got[0].Provider, got[1].Provider}
	// supportedProviders order is aws, azurerm, google — azurerm drops out
	// so aws must come before google.
	if providers[0] != "aws" || providers[1] != "google" {
		t.Errorf("order = %v, want [aws google]", providers)
	}
}

// Test_FormatDeprecated_Text exercises the default pipe-friendly format.
// We check for substrings rather than exact output so minor wording
// tweaks don't force the test to follow; the contract we care about is
// "each entry appears on its own line with provider and kind visible".
func Test_FormatDeprecated_Text(t *testing.T) {
	d := []ProviderDeprecations{
		{
			Provider:    "aws",
			Resources:   map[string]string{"aws_retired": "use aws_new"},
			DataSources: map[string]string{"aws_old_data": ""},
		},
	}

	out, err := FormatDeprecated(d, "text")
	if err != nil {
		t.Fatalf("FormatDeprecated text: %v", err)
	}

	for _, want := range []string{
		"aws:",
		"aws_retired (resource)",
		"use aws_new",
		"aws_old_data (data)",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\ngot:\n%s", want, out)
		}
	}
}

// Test_FormatDeprecated_JSON verifies the JSON output round-trips back
// into the ProviderDeprecations shape so downstream tools can consume
// it without bespoke parsing.
func Test_FormatDeprecated_JSON(t *testing.T) {
	in := []ProviderDeprecations{
		{Provider: "aws", Resources: map[string]string{"aws_retired": "use aws_new"}},
	}

	out, err := FormatDeprecated(in, "json")
	if err != nil {
		t.Fatalf("FormatDeprecated json: %v", err)
	}

	var round []ProviderDeprecations
	if err := json.Unmarshal([]byte(out), &round); err != nil {
		t.Fatalf("json round-trip: %v\npayload: %s", err, out)
	}

	if len(round) != 1 || round[0].Provider != "aws" {
		t.Errorf("round-trip mismatch: got %#v", round)
	}
	if round[0].Resources["aws_retired"] != "use aws_new" {
		t.Errorf("round-trip lost the message: %#v", round[0].Resources)
	}
}

// Test_FormatDeprecated_Empty checks the UX for a clean provider
// snapshot — the user should get a readable "nothing to see" line, not
// an empty string that looks like a silent failure.
func Test_FormatDeprecated_Empty(t *testing.T) {
	out, err := FormatDeprecated(nil, "text")
	if err != nil {
		t.Fatalf("FormatDeprecated: %v", err)
	}
	if !strings.Contains(out, "No deprecated") {
		t.Errorf("empty-case output should say 'No deprecated...'; got %q", out)
	}
}
