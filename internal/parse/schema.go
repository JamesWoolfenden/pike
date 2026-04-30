package parse

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/jameswoolfenden/pike/internal/tfinstall"
	"github.com/rs/zerolog/log"
)

// providerSource maps the short provider names pike has always accepted on
// the CLI (aws, azurerm, google) to their canonical registry sources.
// Stored as a var rather than a switch so tests can substitute mirrors.
var providerSources = map[string]string{
	"aws":     "hashicorp/aws",
	"azurerm": "hashicorp/azurerm",
	"google":  "hashicorp/google",
}

// parseFromSchema extracts a provider's resource and datasource lists by
// running `terraform providers schema -json` against a throwaway workspace
// that has just had the provider initialised.
//
// This is orders of magnitude faster than the historical doc-scanning path
// (one terraform call + one JSON parse vs. walking tens of thousands of
// markdown files and regex-matching each one). It's also authoritative:
// names come from the provider binary's schema, not from doc conventions.
//
// Requires:
//   - terraform (or a tofu binary symlinked as terraform) on PATH, or
//     failing that, network access to install one via hc-install.
//   - network access to the provider registry for `terraform init`.
//
// Callers that can't meet these requirements should supply a codebase
// argument to Parse so the docs-based fallback can run.
func parseFromSchema(name string) (*provider, error) {
	source, ok := providerSources[name]
	if !ok {
		return nil, fmt.Errorf("unsupported provider %q for schema-based parse: known providers are aws, azurerm, google", name)
	}

	tfPath, err := locateTerraform()
	if err != nil {
		return nil, fmt.Errorf("locating terraform: %w", err)
	}

	workDir, err := os.MkdirTemp("", "pike-parse-"+name+"-")
	if err != nil {
		return nil, fmt.Errorf("creating temp workspace: %w", err)
	}

	defer func() {
		// Best-effort cleanup; the temp dir is inside the OS temp tree so
		// a leftover is not dangerous, but it's worth logging so a repeated
		// failure gets noticed rather than silently filling /tmp.
		if rmErr := os.RemoveAll(workDir); rmErr != nil {
			log.Warn().Err(rmErr).Str("dir", workDir).Msg("failed to clean up parse temp workspace")
		}
	}()

	// Minimal config: declare the provider, don't configure it. We don't
	// need valid credentials — `terraform providers schema` only requires
	// the provider plugin to be installed, not to be able to reach its
	// backend.
	cfg := fmt.Sprintf(`terraform {
  required_providers {
    %[1]s = {
      source = %[2]q
    }
  }
}

provider %[1]q {}
`, name, source)

	if err := os.WriteFile(filepath.Join(workDir, "providers.tf"), []byte(cfg), 0o600); err != nil {
		return nil, fmt.Errorf("writing providers.tf: %w", err)
	}

	tf, err := tfexec.NewTerraform(workDir, tfPath)
	if err != nil {
		return nil, fmt.Errorf("tfexec.NewTerraform: %w", err)
	}

	// Provider downloads can be slow on a cold registry or a slow link, but
	// should never hang forever. Five minutes covers a worst-case AWS
	// provider download (~400MB zipped) on a residential connection.
	initCtx, cancelInit := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancelInit()

	log.Info().Str("provider", name).Msg("terraform init (schema path)")
	if err := tf.Init(initCtx, tfexec.Upgrade(false)); err != nil {
		return nil, fmt.Errorf("terraform init for %s: %w", name, err)
	}

	schemaCtx, cancelSchema := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancelSchema()

	schemas, err := tf.ProvidersSchema(schemaCtx)
	if err != nil {
		return nil, fmt.Errorf("terraform providers schema for %s: %w", name, err)
	}

	p := extractMembers(schemas, source)
	p.ProviderVersion = readLockVersion(workDir, source)
	return p, nil
}

// extractMembers pulls resource and datasource names out of the schema map
// returned by terraform-exec. The map key is the fully-qualified registry
// address (e.g. registry.terraform.io/hashicorp/aws) but we only know the
// source suffix (hashicorp/aws), so we match by suffix rather than
// equality.
//
// Deprecation tracking: when a resource/datasource block has Deprecated=true
// we record the block's Description as the deprecation message. Providers
// conventionally prefix the description with "Deprecated: use X instead" so
// the raw description is serviceable; if a provider leaves Description empty
// we still record the entry with an empty-string value so consumers can
// treat presence-of-key as the deprecation signal.
func extractMembers(schemas *tfjson.ProviderSchemas, source string) *provider {
	out := &provider{}
	for key, ps := range schemas.Schemas {
		if !matchesSource(key, source) {
			continue
		}
		for name, sch := range ps.ResourceSchemas {
			out.Resources = append(out.Resources, name)
			if sch != nil && sch.Block != nil && sch.Block.Deprecated {
				if out.DeprecatedResources == nil {
					out.DeprecatedResources = map[string]string{}
				}
				out.DeprecatedResources[name] = sch.Block.Description
			}
		}
		for name, sch := range ps.DataSourceSchemas {
			out.DataSources = append(out.DataSources, name)
			if sch != nil && sch.Block != nil && sch.Block.Deprecated {
				if out.DeprecatedData == nil {
					out.DeprecatedData = map[string]string{}
				}
				out.DeprecatedData[name] = sch.Block.Description
			}
		}
		break
	}
	sort.Strings(out.Resources)
	sort.Strings(out.DataSources)
	return out
}

// matchesSource reports whether a fully-qualified schema key references the
// given short source. Accepts both the hostnameless form ("hashicorp/aws")
// and the hostname'd form ("registry.terraform.io/hashicorp/aws").
func matchesSource(key, source string) bool {
	if key == source {
		return true
	}
	return strings.HasSuffix(key, "/"+source)
}

func locateTerraform() (string, error) {
	return tfinstall.LocateTerraform()
}

// readLockVersion extracts the installed provider version from the
// .terraform.lock.hcl written by `terraform init`. The lock file format is:
//
//	provider "registry.terraform.io/hashicorp/aws" {
//	  version   = "5.82.2"
//	}
//
// source is the short registry path (e.g. "hashicorp/aws"); we match on the
// trailing segment so we find the right block regardless of whether the full
// hostname prefix is present. Returns "" on any read or parse error so the
// caller treats a missing version as non-fatal.
func readLockVersion(workDir, source string) string {
	data, err := os.ReadFile(filepath.Join(workDir, ".terraform.lock.hcl"))
	if err != nil {
		return ""
	}
	// Match only on the name component (e.g. "aws") to handle both
	// "hashicorp/aws" and "registry.terraform.io/hashicorp/aws".
	name := source[strings.LastIndex(source, "/")+1:]
	lines := strings.Split(string(data), "\n")
	inBlock := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "/"+name+`"`) {
			inBlock = true
			continue
		}
		if inBlock && strings.Contains(trimmed, "version") && strings.Contains(trimmed, "=") {
			// version   = "5.82.2"
			parts := strings.SplitN(trimmed, `"`, 3)
			if len(parts) >= 2 {
				return parts[1]
			}
		}
		if inBlock && trimmed == "}" {
			break
		}
	}
	return ""
}
