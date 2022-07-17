package pike

import "github.com/hashicorp/hcl/hcl/ast"

type template struct {
	Resource Resource `json:"resource"`
	API      string   `json:"api"`
	Provider string   `json:"provider"`
}

// Resource object for the HCL parser
type Resource struct {
	name string
	path string
	code ast.ObjectItem
}
