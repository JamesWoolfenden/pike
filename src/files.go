package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/aws/resource/aws_lb.json
var awsLb []byte

//go:embed mapping/aws/resource/aws_s3_bucket.json
var awsS3Bucket []byte

//go:embed mapping/aws/resource/aws_s3_bucket_acl.json
var awsS3BucketACL []byte

//go:embed mapping/aws/resource/aws_s3_bucket_versioning.json
var awsS3BucketVersioning []byte

//go:embed mapping/aws/resource/aws_s3_bucket_server_side_encryption_configuration.json
var awsS3BucketServerSideEncryptionConfiguration []byte

//go:embed mapping/aws/resource/aws_s3_bucket_public_access_block.json
var awsS3BucketPublicAccessBlock []byte

//go:embed mapping/aws/resource/aws_instance.json
var awsInstance []byte

//go:embed mapping/aws/resource/aws_security_group.json
var awsSecurityGroup []byte

//go:embed mapping/aws/resource/aws_security_group_rule.json
var awsSecurityGroupRule []byte

//go:embed mapping/aws/resource/aws_lambda_function.json
var awsLambdaFunction []byte

//go:embed mapping/aws/resource/aws_vpc.json
var awsVpc []byte

//go:embed mapping/aws/resource/aws_subnet.json
var awsSubnet []byte

//go:embed mapping/aws/resource/aws_network_acl.json
var awsNetworkACL []byte

//go:embed mapping/aws/resource/aws_kms_key.json
var awsKmsKey []byte

//go:embed mapping/aws/resource/aws_iam_role.json
var awsIamRole []byte

//go:embed mapping/aws/resource/aws_iam_role_policy.json
var awsIamRolePolicy []byte

//go:embed mapping/aws/resource/aws_iam_role_policy_attachment.json
var awsIamRolePolicyAttachment []byte

//go:embed mapping/aws/resource/aws_iam_policy.json
var awsIamPolicy []byte

//go:embed mapping/aws/resource/aws_iam_instance_profile.json
var awsIamInstanceProfile []byte

//go:embed mapping/aws/resource/aws_iam_access_key.json
var awsIamAccessKey []byte

//go:embed mapping/aws/resource/aws_iam_group.json
var awsIamGroup []byte

//go:embed mapping/aws/resource/aws_iam_group_membership.json
var awsIamGroupMembership []byte

//go:embed mapping/aws/resource/aws_iam_group_policy.json
var awsIamGroupPolicy []byte

//go:embed mapping/aws/resource/aws_iam_group_policy_attachment.json
var awsIamGroupPolicyAttachment []byte

//go:embed mapping/aws/resource/aws_iam_policy_attachment.json
var awsIamPolicyAttachment []byte

//go:embed mapping/aws/resource/aws_iam_service_linked_role.json
var awsIamServiceLinkedRole []byte

//go:embed mapping/aws/resource/aws_iam_user.json
var awsIamUser []byte

//go:embed mapping/aws/resource/aws_iam_user_login_profile.json
var awsIamUserLoginProfile []byte

//go:embed mapping/aws/resource/aws_iam_user_policy.json
var awsIamUserPolicy []byte

//go:embed mapping/aws/resource/aws_iam_user_policy_attachment.json
var awsIamUserPolicyAttachment []byte

//go:embed mapping/aws/resource/aws_mq_broker.json
var awsMqBroker []byte

//go:embed mapping/aws/resource/aws_mq_configuration.json
var awsMqConfiguration []byte

//go:embed mapping/aws/resource/aws_cloudwatch_log_group.json
var awsCloudwatchLogGroup []byte

//go:embed mapping/aws/resource/aws_cloudwatch_event_rule.json
var awsCloudwatchEventRule []byte

//go:embed mapping/aws/resource/aws_cloudwatch_log_metric_filter.json
var awsCloudwatchLogMetricFilter []byte

//go:embed mapping/aws/resource/aws_cloudwatch_log_resource_policy.json
var awsCloudwatchLogResourcePolicy []byte

//go:embed mapping/aws/resource/aws_cloudwatch_log_subscription_filter.json
var awsCloudwatchLogSubscriptionFilter []byte

//go:embed mapping/aws/resource/aws_cloudwatch_metric_alarm.json
var awsCloudwatchMetricAlarm []byte

//go:embed mapping/aws/resource/aws_cloudwatch_event_target.json
var awsCloudwatchEventTarget []byte

//go:embed mapping/aws/resource/aws_route53_record.json
var awsRoute53Record []byte

//go:embed mapping/aws/resource/aws_sns_topic.json
var awsSnsTopic []byte

//go:embed mapping/aws/resource/aws_key_pair.json
var awsKeyPair []byte

//go:embed mapping/aws/resource/aws_db_instance.json
var awsDbInstance []byte

//go:embed mapping/aws/resource/aws_dynamodb_table.json
var awsDynamodbTable []byte

//go:embed mapping/aws/resource/aws_ssm_parameter.json
var awsSsmParameter []byte

//go:embed mapping/aws/resource/aws_route.json
var awsRoute []byte

//go:embed mapping/gcp/google_compute_instance.json
var googleComputeInstance []byte
