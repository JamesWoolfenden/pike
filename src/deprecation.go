package pike

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/jameswoolfenden/pike/parse"
	"github.com/rs/zerolog/log"
)

// deprecationEntry mirrors the shape written by parse/parse.go's provider
// struct (only the fields this package needs). The JSON tags MUST stay in
// sync with that struct; see parse/parse.go for why the dataSources tag is
// camelCase.
type deprecationEntry struct {
	DeprecatedResources map[string]string `json:"deprecatedResources"`
	DeprecatedData      map[string]string `json:"deprecatedData"`
}

// deprecationStore caches the per-provider deprecation maps extracted from
// parse.EmbeddedMembers. It's populated on first lookup per provider so
// pike pays the JSON-decode cost only when scan actually encounters a
// resource that triggers a check — a user scanning a single aws_s3_bucket
// never parses the azure or gcp members.
//
// Concurrent scanners (e.g. the CI runner scanning multiple projects)
// share a process-wide cache; sync.Once per provider keeps the decode
// single-shot.
type deprecationStore struct {
	once    map[string]*sync.Once
	entries map[string]deprecationEntry
	mu      sync.Mutex
}

var deprecations = &deprecationStore{
	once:    map[string]*sync.Once{},
	entries: map[string]deprecationEntry{},
}

// seedDeprecationForTest is a package-private test hook. It installs a
// precomputed deprecationEntry for a given provider and marks its
// sync.Once as consumed so loadEntry won't overwrite it from embedded
// bytes. Tests should restore prior state via t.Cleanup.
//
// Kept in the production file (rather than _test.go) so the symbol is
// visible to any test in the package without an export_test.go dance,
// but it's unexported so external callers can't reach it.
func seedDeprecationForTest(provider string, entry deprecationEntry) {
	deprecations.mu.Lock()
	defer deprecations.mu.Unlock()
	once := &sync.Once{}
	once.Do(func() {}) // consume so loadEntry is a no-op
	deprecations.once[provider] = once
	deprecations.entries[provider] = entry
}

// clearDeprecationsForTest wipes the cache so the next lookup goes
// through the real embedded-bytes path. Useful in t.Cleanup.
func clearDeprecationsForTest() {
	deprecations.mu.Lock()
	defer deprecations.mu.Unlock()
	deprecations.once = map[string]*sync.Once{}
	deprecations.entries = map[string]deprecationEntry{}
}

// normaliseProvider maps the various aliases pike accepts (gcp/google,
// azure/azurerm) to the canonical short name used as the embed key.
func normaliseProvider(p string) string {
	switch strings.ToLower(p) {
	case "aws":
		return "aws"
	case "azure", "azurerm":
		return "azurerm"
	case "gcp", "google":
		return "google"
	default:
		return ""
	}
}

// loadEntry decodes the embedded members JSON for a provider the first
// time it's asked for, then returns the cached copy. A missing or
// malformed embed is logged once and treated as "nothing deprecated" —
// pike's scan should not hard-fail on advisory metadata.
func (s *deprecationStore) loadEntry(provider string) deprecationEntry {
	key := normaliseProvider(provider)
	if key == "" {
		return deprecationEntry{}
	}

	s.mu.Lock()
	if _, ok := s.once[key]; !ok {
		s.once[key] = &sync.Once{}
	}
	once := s.once[key]
	s.mu.Unlock()

	once.Do(func() {
		raw := parse.EmbeddedMembers(key)
		if len(raw) == 0 {
			log.Debug().Str("provider", key).Msg("no embedded members for provider; deprecation lookup disabled")
			return
		}
		var entry deprecationEntry
		if err := json.Unmarshal(raw, &entry); err != nil {
			log.Warn().Err(err).Str("provider", key).Msg("failed to decode embedded members; deprecation lookup disabled for this provider")
			return
		}
		s.mu.Lock()
		s.entries[key] = entry
		s.mu.Unlock()
	})

	s.mu.Lock()
	defer s.mu.Unlock()
	return s.entries[key]
}

// DeprecationMessage reports whether a given terraform resource or
// datasource is marked deprecated by its provider. ok is true when the
// provider's latest schema has flagged it; msg is the provider-supplied
// deprecation description (often "Deprecated: use X instead"), which may
// be empty even when ok is true if the provider didn't include one.
//
// isData distinguishes a `data "foo" {}` block from a `resource "foo" {}`
// block: providers sometimes deprecate one without the other, so the
// caller must pass the correct kind.
func DeprecationMessage(provider, name string, isData bool) (msg string, ok bool) {
	entry := deprecations.loadEntry(provider)
	if isData {
		msg, ok = entry.DeprecatedData[name]
	} else {
		msg, ok = entry.DeprecatedResources[name]
	}
	return msg, ok
}

// warnIfDeprecated emits a single log.Warn per (provider, resource-type,
// kind) tuple encountered during a scan. The `warned` set is shared
// across the caller's loop so the same resource declared multiple times
// in a user's HCL only produces one line of output.
//
// ResourceV2 field naming is historical and slightly counter-intuitive:
//   - TypeName holds the HCL block type ("resource" or "data")
//   - Name holds the terraform resource type (e.g. "aws_s3_bucket")
//
// We use Name as the lookup key against the provider's deprecation map.
func warnIfDeprecated(r ResourceV2, warned map[string]struct{}) {
	if r.Name == "" || r.Provider == "" {
		return
	}
	isData := r.TypeName == "data"
	msg, ok := DeprecationMessage(r.Provider, r.Name, isData)
	if !ok {
		return
	}
	key := r.Provider + "\x00" + r.Name
	if isData {
		key += "\x00data"
	}
	if _, already := warned[key]; already {
		return
	}
	warned[key] = struct{}{}

	evt := log.Warn().Str("provider", r.Provider).Str("resource", r.Name)
	if isData {
		evt = evt.Str("kind", "data")
	}
	if msg != "" {
		evt = evt.Str("message", msg)
	}
	evt.Msg("deprecated in latest provider schema")
}

// supportedProviders is the canonical list the deprecated command walks
// when no specific provider is requested. Kept deliberately in sync with
// parse.EmbeddedMembers — if you add a provider there, add it here.
var supportedProviders = []string{"aws", "azurerm", "google"}

// ProviderDeprecations is the serialisable shape returned by Deprecated().
// Kinds are split because the same name can appear as both a resource and
// a datasource in the same provider, and conflating them loses detail.
type ProviderDeprecations struct {
	Provider    string            `json:"provider"`
	Resources   map[string]string `json:"resources,omitempty"`
	DataSources map[string]string `json:"dataSources,omitempty"`
}

// Deprecated returns the known-deprecated resources/datasources for one
// provider (aws/azurerm/google, with gcp/azure aliases accepted) or for
// all supported providers when called with the empty string. Providers
// with nothing deprecated are omitted from the result so callers can
// cleanly skip over them.
//
// This is the data source for the `pike deprecated` CLI. It reads from
// the same lazily-cached store that scan-time warnings use, so running
// `pike deprecated` after a scan is free.
func Deprecated(provider string) []ProviderDeprecations {
	var names []string
	if provider == "" {
		names = supportedProviders
	} else {
		n := normaliseProvider(provider)
		if n == "" {
			return nil
		}
		names = []string{n}
	}

	out := make([]ProviderDeprecations, 0, len(names))
	for _, n := range names {
		entry := deprecations.loadEntry(n)
		if len(entry.DeprecatedResources) == 0 && len(entry.DeprecatedData) == 0 {
			continue
		}
		out = append(out, ProviderDeprecations{
			Provider:    n,
			Resources:   entry.DeprecatedResources,
			DataSources: entry.DeprecatedData,
		})
	}
	return out
}

// FormatDeprecated renders the Deprecated() output. format is one of
// "text" (default, grep-friendly) or "json" (pretty-printed for
// downstream tooling). Unknown formats fall back to text rather than
// erroring — the CLI surface is informational, not strict.
//
// Text format:
//
//	aws:
//	  aws_retired_thing (resource) — Deprecated: use aws_new_thing
//	  aws_old_data (data) — ...
//	azurerm:
//	  ...
//
// An empty result (nothing deprecated across all requested providers)
// returns a single-line notice rather than an empty string so users
// don't mistake success for "the command silently did nothing".
func FormatDeprecated(d []ProviderDeprecations, format string) (string, error) {
	if len(d) == 0 {
		return "No deprecated resources or datasources found in the embedded provider schemas.\n", nil
	}

	switch strings.ToLower(format) {
	case "json":
		b, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			return "", fmt.Errorf("marshalling deprecation data: %w", err)
		}
		return string(b) + "\n", nil
	default:
		var b strings.Builder
		for _, p := range d {
			fmt.Fprintf(&b, "%s:\n", p.Provider)

			for _, name := range sortedKeys(p.Resources) {
				writeDeprecationLine(&b, name, "resource", p.Resources[name])
			}
			for _, name := range sortedKeys(p.DataSources) {
				writeDeprecationLine(&b, name, "data", p.DataSources[name])
			}
		}
		return b.String(), nil
	}
}

func sortedKeys(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func writeDeprecationLine(b *strings.Builder, name, kind, msg string) {
	// Single-line output keeps the command pipe-friendly
	// (`pike deprecated | grep azurerm_blob`).
	msg = strings.ReplaceAll(msg, "\n", " ")
	if msg == "" {
		fmt.Fprintf(b, "  %s (%s)\n", name, kind)
	} else {
		fmt.Fprintf(b, "  %s (%s) — %s\n", name, kind, msg)
	}
}
