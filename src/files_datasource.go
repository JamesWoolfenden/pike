package pike

import (
	_ "embed" //required for embed
)

//go:embed mapping/aws/data/aws_vpcs.json
var data_aws_vpcs []byte

//go:embed mapping/aws/data/aws_subnets_ids.json
var data_aws_subnet_ids []byte
