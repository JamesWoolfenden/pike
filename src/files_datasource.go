package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/aws/data/template.json
var placeholder []byte

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

//go:embed mapping/aws/data/aws_s3_bucket.json
var dataAwsS3Bucket []byte

//go:embed mapping/aws/data/aws_inspector_rules_packages.json
var dataAwsInspectorRulesPackages []byte

//go:embed mapping/aws/data/aws_route53_zone.json
var dataAwsRoute53Zone []byte

//go:embed mapping/aws/data/aws_kms_ciphertext.json
var dataAwsKmsCiphertext []byte
