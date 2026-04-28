package pike

import (
	"testing"
)

func TestRuntime_UnsupportedProvider(t *testing.T) {
	t.Parallel()

	err := Runtime(".", "terraform", nil, false, "aws")
	if err == nil {
		t.Fatal("Runtime() with unsupported provider expected error, got nil")
	}

	var unsupported *unsupportedRuntimeProviderError
	found := false
	if e, ok := err.(*unsupportedRuntimeProviderError); ok {
		unsupported = e
		found = true
	}
	if !found {
		t.Fatalf("Runtime() error type = %T, want *unsupportedRuntimeProviderError", err)
	}
	if unsupported.provider != "aws" {
		t.Errorf("unsupportedRuntimeProviderError.provider = %q, want %q", unsupported.provider, "aws")
	}
}

func TestRuntime_AzureUnsupported(t *testing.T) {
	t.Parallel()

	if err := Runtime(".", "terraform", nil, false, "azure"); err == nil {
		t.Fatal("Runtime() with azure provider expected error, got nil")
	}
}

func TestRuntime_EmptyLocation(t *testing.T) {
	t.Parallel()

	err := Runtime("", "terraform", nil, false, "")
	if err == nil {
		t.Fatal("Runtime() with empty dir and nil file expected error, got nil")
	}
}
