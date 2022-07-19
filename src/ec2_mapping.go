package pike

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed ec2.json
var ec2raw []byte

// EC2Test Anonymous parsing
func EC2Test() map[string]interface{} {
	var mappings []interface{}
	err := json.Unmarshal(ec2raw, &mappings)
	if err != nil {
		log.Print(err)
	}
	return mappings[0].(map[string]interface{})
}
