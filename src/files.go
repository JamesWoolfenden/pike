package pike

import (
	_ "embed" //required for embed
)

//go:embed mapping/aws/resource/aws_s3_bucket.json
var aws_s3_bucket []byte

//go:embed mapping/aws/resource/aws_instance.json
var aws_instance []byte

//go:embed mapping/aws/resource/aws_security_group.json
var aws_security_group []byte

//go:embed mapping/aws/resource/aws_lambda_function.json
var aws_lambda_function []byte

//go:embed mapping/aws/resource/aws_vpc.json
var aws_vpc []byte

//go:embed mapping/aws/resource/aws_subnet.json
var aws_subnet []byte

//go:embed mapping/aws/resource/aws_network_acl.json
var aws_network_acl []byte

//go:embed mapping/aws/resource/aws_kms_key.json
var aws_kms_key []byte

//go:embed mapping/aws/resource/aws_iam_role.json
var aws_iam_role []byte

//go:embed mapping/aws/resource/aws_mq_broker.json
var aws_mq_broker []byte

//go:embed mapping/aws/resource/aws_mq_configuration.json
var aws_mq_configuration []byte

//go:embed mapping/gcp/google_compute_instance.json
var google_compute_instance []byte
