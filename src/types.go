package pike

import "github.com/hashicorp/hcl/hcl/ast"

type template struct {
	Resource Resource `json:"resource"`
	Provider string   `json:"provider"`
	Template string   `json:"template"`
}

// Resource object for the HCL parser
type Resource struct {
	name string
	path string
	code ast.ObjectItem
}

// Sorted is to help split out permission to the relevant auth
type Sorted struct {
	AWS   []string
	GCP   []string
	AZURE []string
}
