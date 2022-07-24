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

//go:embed mapping/aws/aws_lambda.json
var aws_lambda []byte