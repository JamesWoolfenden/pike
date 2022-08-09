package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/aws/data/aws_vpcs.json
var dataAwsVpcs []byte

//go:embed mapping/aws/data/aws_subnets_ids.json
var dataAwsSubnetIds []byte

//go:embed mapping/aws/data/aws_ami.json
var dataAwsAmi []byte

//go:embed mapping/aws/data/aws_availability_zones.json
var dataAwsAvailabilityZones []byte

//go:embed mapping/aws/data/aws_iam_policy.json
var dataAwsIamPolicy []byte

//go:embed mapping/aws/data/aws_iam_role.json
var dataAwsIamRole []byte

//go:embed mapping/aws/data/aws_vpc.json
var dataAwsVpc []byte
