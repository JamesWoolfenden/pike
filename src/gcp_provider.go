package pike

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
)

// hasGoogleUserProjectOverride detects quota-project permission requirements,
// independent of the resource mappings. billing_project may be an expression
// or omitted (the provider then uses the resource project), so its literal
// value is not needed to infer serviceusage.services.use.
func hasGoogleUserProjectOverride(block *hclsyntax.Block) bool {
	if block.Type != "provider" || len(block.Labels) != 1 {
		return false
	}
	name := strings.ToLower(block.Labels[0])
	if name != "google" && name != "google-beta" {
		return false
	}
	attribute, ok := block.Body.Attributes["user_project_override"]
	if !ok {
		return false // The provider defaults to false.
	}
	value, diags := attribute.Expr.Value(nil)
	if diags.HasErrors() || !value.IsKnown() {
		// Variables/locals cannot be resolved by this static scan. Include the
		// permission conservatively rather than risk under-reporting access.
		return true
	}
	if value.IsNull() {
		return false
	}
	boolean, err := convert.Convert(value, cty.Bool)
	return err == nil && boolean.True()
}
