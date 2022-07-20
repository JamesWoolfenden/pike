package pike

import (
	_ "embed" //required for embed
)

//go:embed s3_bucket.json
var s3 []byte

//go:embed ec2.json
var ec2raw []byte
