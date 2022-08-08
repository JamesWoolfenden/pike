package pike

import (
	_ "embed" //required for embed
)

//go:embed mapping/aws/resource/aws_s3_bucket.json
var aws_s3_bucket []byte

//go:embed mapping/aws/resource/aws_s3_bucket_acl.json
var aws_s3_bucket_acl []byte

//go:embed mapping/aws/resource/aws_s3_bucket_versioning.json
var aws_s3_bucket_versioning []byte

//go:embed mapping/aws/resource/aws_s3_bucket_server_side_encryption_configuration.json
var aws_s3_bucket_server_side_encryption_configuration []byte

//go:embed mapping/aws/resource/aws_s3_bucket_public_access_block.json
var aws_s3_bucket_public_access_block []byte

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

//go:embed mapping/aws/resource/aws_iam_role_policy.json
var aws_iam_role_policy []byte

//go:embed mapping/aws/resource/aws_iam_role_policy_attachment.json
var aws_iam_role_policy_attachment []byte

//go:embed mapping/aws/resource/aws_iam_policy.json
var aws_iam_policy []byte

//go:embed mapping/aws/resource/aws_iam_instance_profile.json
var aws_iam_instance_profile []byte

//go:embed mapping/aws/resource/aws_mq_broker.json
var aws_mq_broker []byte

//go:embed mapping/aws/resource/aws_mq_configuration.json
var aws_mq_configuration []byte

//go:embed mapping/aws/resource/aws_cloudwatch_log_group.json
var aws_cloudwatch_log_group []byte

//go:embed mapping/aws/resource/aws_route53_record.json
var aws_route53_record []byte

//go:embed mapping/aws/resource/aws_sns_topic.json
var aws_sns_topic []byte

//go:embed mapping/aws/resource/aws_key_pair.json
var aws_key_pair []byte

//go:embed mapping/aws/resource/aws_db_instance.json
var aws_db_instance []byte

//go:embed mapping/aws/resource/aws_dynamodb_table.json
var aws_dynamodb_table []byte

//go:embed mapping/gcp/google_compute_instance.json
var google_compute_instance []byte
