package pike

import (
	_ "embed" //required for embed
)

//go:embed mapping/aws/aws_s3_bucket.json
var aws_s3_bucket []byte

//go:embed mapping/aws/aws_instance.json
var aws_instance []byte

//go:embed mapping/aws/aws_security_group.json
var aws_security_group []byte

//go:embed mapping/aws/aws_lambda_function.json
var aws_lambda_function []byte

//go:embed mapping/aws/aws_vpc.json
var aws_vpc []byte

//go:embed mapping/aws/aws_subnet.json
var aws_subnet []byte

//go:embed mapping/aws/aws_network_acl.json
var aws_network_acl []byte

//go:embed mapping/gcp/google_compute_instance.json
var google_compute_instance []byte
