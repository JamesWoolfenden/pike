package pike

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/assert"
)

func parseHCL(t *testing.T, src string) *hclsyntax.Body {
	t.Helper()
	f, diags := hclparse.NewParser().ParseHCL([]byte(src), "test.tf")
	if diags.HasErrors() {
		t.Fatalf("parse: %v", diags)
	}
	return f.Body.(*hclsyntax.Body)
}

func TestHasAWSDefaultTags(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"with default_tags", "provider \"aws\" {\n  default_tags {\n    tags = { env = \"prod\" }\n  }\n}\n", true},
		{"aws no default_tags", "provider \"aws\" { region = \"eu-west-2\" }\n", false},
		{"non-aws provider", "provider \"google\" {\n  default_tags {}\n}\n", false},
		{"no provider", "resource \"aws_instance\" \"x\" {}\n", false},
		{"mixed-case label", "provider \"AWS\" {\n  default_tags {}\n}\n", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, HasAWSDefaultTags(parseHCL(t, tt.src)))
		})
	}
	assert.False(t, HasAWSDefaultTags(nil))
}

func Test_fileStringEmptyError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Invoke", "no file provided"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &fileStringEmptyError{}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}

func Test_isIAMBindingResource(t *testing.T) {
	tests := []struct {
		name         string
		resourceType string
		want         bool
	}{
		{
			name:         "google iam member",
			resourceType: "google_storage_bucket_iam_member",
			want:         true,
		},
		{
			name:         "google iam binding",
			resourceType: "google_storage_bucket_iam_binding",
			want:         true,
		},
		{
			name:         "google compute instance iam member",
			resourceType: "google_compute_instance_iam_member",
			want:         true,
		},
		{
			name:         "aws iam member - not google",
			resourceType: "aws_iam_member",
			want:         false,
		},
		{
			name:         "google regular resource",
			resourceType: "google_storage_bucket",
			want:         false,
		},
		{
			name:         "azurerm iam member",
			resourceType: "azurerm_role_assignment_iam_member",
			want:         false,
		},
		{
			name:         "empty string",
			resourceType: "",
			want:         false,
		},
		{
			name:         "just iam_member",
			resourceType: "iam_member",
			want:         false,
		},
		{
			name:         "just iam_binding",
			resourceType: "iam_binding",
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isIAMBindingResource(tt.resourceType)
			if got != tt.want {
				t.Errorf("isIAMBindingResource(%q) = %v, want %v", tt.resourceType, got, tt.want)
			}
		})
	}
}
