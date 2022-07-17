package ec2

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed ec2.json
var ec2raw []byte

// Mapping Struct
type Mapping struct {
	Apply   []Permission `json:"apply"`
	Destroy []Permission `json:"destroy"`
	Plan    []Permission `json:"plan"`
}

// Permission Object
type Permission struct {
	Attribute  string `json:"attribute"`
	Permission string `json:"permission"`
}

func Test() {
	var mappings []Mapping
	_ = json.Unmarshal(ec2raw, &mappings)

	fmt.Println(mappings)
}
