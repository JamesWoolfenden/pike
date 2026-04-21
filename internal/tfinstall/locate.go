// Package tfinstall locates or installs a terraform/tofu binary. It is
// shared by src/ (the scan path) and parse/ (the schema extraction path)
// so that the binary-search and hc-install fallback live in exactly one place.
package tfinstall

import (
	"context"
	"fmt"
	"os/exec"
	"sync"

	goversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/rs/zerolog/log"
)

// Version is the Terraform version fetched via hc-install when no binary is
// found on PATH. Kept as a single constant here so both callers stay in sync.
const Version = "1.5.4"

var mu sync.Mutex

// LocateTerraform returns the path to a tofu or terraform binary, preferring
// tofu. If neither is on PATH it installs Terraform Version via hc-install.
func LocateTerraform() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, bin := range []string{"tofu", "terraform"} {
		if p, err := exec.LookPath(bin); err == nil && p != "" {
			return p, nil
		}
	}

	log.Info().Msgf("no tofu/terraform on PATH; installing terraform %s", Version)

	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: goversion.Must(goversion.NewVersion(Version)),
	}

	p, err := installer.Install(context.Background())
	if err != nil {
		return "", fmt.Errorf("installing terraform %s: %w", Version, err)
	}

	return p, nil
}
