package pike

import (
	"encoding/json"
	"log"
)

// GetAWSPermissions for AWS resources
func GetAWSPermissions(result ResourceV2) []string {

	var Permissions []string
	if result.TypeName == "resource" {
		Permissions = GetAWSResourcePermissions(result)
	} else {
		Permissions = GetAWSDataPermissions(result)
	}

	return Permissions
}

// GetAWSResourcePermissions looks up permissions required for resources
func GetAWSResourcePermissions(result ResourceV2) []string {

	TFLookup := map[string]interface{}{
		"aws_s3_bucket":            awsS3Bucket,
		"aws_s3_bucket_acl":        awsS3BucketACL,
		"aws_s3_bucket_versioning": awsS3BucketVersioning,
		"aws_s3_bucket_server_side_encryption_configuration": awsS3BucketServerSideEncryptionConfiguration,
		"aws_s3_bucket_public_access_block":                  awsS3BucketPublicAccessBlock,
		"aws_s3_bucket_logging":                              awsS3BucketLogging,
		"aws_s3_bucket_lifecycle_configuration":              awsS3BucketLifecycleConfiguration,
		"aws_s3_bucket_policy":                               awsS3BucketPolicy,
		"aws_s3_bucket_object":                               awsS3Object,
		"aws_s3_object":                                      awsS3Object,
		"aws_instance":                                       awsInstance,
		"aws_security_group":                                 awsSecurityGroup,
		"aws_security_group_rule":                            awsSecurityGroupRule,
		"aws_lambda_function":                                awsLambdaFunction,
		"aws_lambda_alias":                                   awsLambdaAlias,
		"aws_lambda_permission":                              awsLambdaPermission,
		"aws_vpc":                                            awsVpc,
		"aws_subnet":                                         awsSubnet,
		"aws_network_acl":                                    awsNetworkACL,
		"aws_kms_key":                                        awsKmsKey,
		"aws_iam_role":                                       awsIamRole,
		"aws_iam_role_policy":                                awsIamRolePolicy,
		"aws_iam_role_policy_attachment":                     awsIamRolePolicyAttachment,
		"aws_iam_policy":                                     awsIamPolicy,
		"aws_iam_instance_profile":                           awsIamInstanceProfile,
		"aws_iam_access_key":                                 awsIamAccessKey,
		"aws_iam_group":                                      awsIamGroup,
		"aws_iam_group_membership":                           awsIamGroupMembership,
		"aws_iam_group_policy":                               awsIamGroupPolicy,
		"aws_iam_group_policy_attachment":                    awsIamGroupPolicyAttachment,
		"aws_iam_policy_attachment":                          awsIamPolicyAttachment,
		"aws_iam_service_linked_role":                        awsIamServiceLinkedRole,
		"aws_iam_user":                                       awsIamUser,
		"aws_iam_user_login_profile":                         awsIamUserLoginProfile,
		"aws_iam_user_policy":                                awsIamUserPolicy,
		"aws_iam_user_policy_attachment":                     awsIamUserPolicyAttachment,
		"aws_mq_broker":                                      awsMqBroker,
		"aws_mq_configuration":                               awsMqConfiguration,
		"aws_cloudwatch_log_group":                           awsCloudwatchLogGroup,
		"aws_cloudwatch_event_rule":                          awsCloudwatchEventRule,
		"aws_cloudwatch_event_target":                        awsCloudwatchEventTarget,
		"aws_cloudwatch_log_metric_filter":                   awsCloudwatchLogMetricFilter,
		"aws_cloudwatch_log_resource_policy":                 awsCloudwatchLogResourcePolicy,
		"aws_cloudwatch_log_subscription_filter":             awsCloudwatchLogSubscriptionFilter,
		"aws_cloudwatch_metric_alarm":                        awsCloudwatchMetricAlarm,
		"aws_route53_record":                                 awsRoute53Record,
		"aws_route53_zone":                                   awsRoute53Zone,
		"aws_sns_topic":                                      awsSnsTopic,
		"aws_sns_topic_subscription":                         awsSnsTopicSubscription,
		"aws_sns_topic_policy":                               awsSnsTopicPolicy,
		"aws_key_pair":                                       awsKeyPair,
		"aws_db_instance":                                    awsDbInstance,
		"aws_dynamodb_table":                                 awsDynamodbTable,
		"aws_ssm_parameter":                                  awsSsmParameter,
		"aws_route":                                          awsRoute,
		"aws_lb":                                             awsLb,
		"aws_alb":                                            awsLb,
		"aws_alb_listener":                                   awsLbListener,
		"aws_lb_listener":                                    awsLbListener,
		"aws_lb_target_group":                                awsLbTargetGroup,
		"aws_alb_target_group":                               awsLbTargetGroup,
		"aws_alb_target_group_attachment":                    awsLbTargetGroupAttachment,
		"aws_lb_target_group_attachment":                     awsLbTargetGroupAttachment,
		"aws_default_security_group":                         awsDefaultSecurityGroup,
		"aws_db_subnet_group":                                awsDbSubnetGroup,
		"aws_wafv2_web_acl":                                  awsWafv2WebACL,
		"aws_apigatewayv2_api":                               awsApigatewayv2Api,
		"aws_api_gateway_rest_api":                           awsAPIGatewayRestAPI,
		"aws_api_gateway_api_key":                            awsApigatewayv2Api,
		"aws_api_gateway_deployment":                         awsApigatewayv2Api,
		"aws_api_gateway_stage":                              awsApigatewayv2Api,
		"aws_api_gateway_integration":                        awsApigatewayv2Api,
		"aws_api_gateway_resource":                           awsApigatewayv2Api,
		"aws_api_gateway_method":                             awsApigatewayv2Api,
		"aws_api_gateway_method_settings":                    awsApigatewayv2Api,
		"aws_api_gateway_method_response":                    awsApigatewayv2Api,
		"aws_api_gateway_integration_response":               awsApigatewayv2Api,
		"aws_api_gateway_usage_plan":                         awsApigatewayv2Api,
		"aws_api_gateway_usage_plan_key":                     awsApigatewayv2Api,
		"aws_api_gateway_account":                            awsAPIGatewayAccount,
		"aws_sqs_queue":                                      awsSqsQueue,
		"aws_sqs_queue_policy":                               awsSqsQueuePolicy,
		"aws_ebs_volume":                                     awsEbsVolume,
		"aws_autoscaling_group":                              awsAutoscalingGroup,
		"aws_autoscaling_attachment":                         awsAutoscalingAttachment,
		"aws_elb":                                            awsElb,
		"aws_internet_gateway":                               awsInternetGateway,
		"aws_launch_configuration":                           awsLaunchConfiguration,
		"aws_ec2_capacity_reservation":                       awsEc2CapacityReservation,
		"aws_network_interface":                              awsNetworkInterface,
		"aws_placement_group":                                awsPlacementGroup,
		"aws_spot_instance_request":                          awsSpotInstanceRequest,
		"aws_volume_attachment":                              awsVolumeAttachment,
		"aws_budgets_budget":                                 awsBudgetsBudget,
		"aws_eip":                                            awsEip,
		"aws_kinesis_firehose_delivery_stream":               awsKinesisFirehoseDeliveryStream,
		"aws_kinesis_stream":                                 awsKinesisStream,
		"aws_kinesis_video_stream":                           awsKinesisVideoStream,
		"aws_elastic_beanstalk_application":                  awsElasticBeanstalkApplication,
		"aws_flow_log":                                       awsFlowLog,
	}

	var Permissions []string

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("%s not implemented", result.Name)
	}

	return Permissions
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetPermissionMap Anonymous parsing
func GetPermissionMap(raw []byte, attributes []string) []string {
	var mappings []interface{}
	err := json.Unmarshal(raw, &mappings)
	if err != nil {
		log.Print(err)
	}
	temp := mappings[0].(map[string]interface{})
	myAttributes := temp["attributes"].(map[string]interface{})

	var found []string

	for _, attribute := range attributes {
		if myAttributes[attribute] != nil {
			entries := myAttributes[attribute].([]interface{})
			for _, entry := range entries {
				found = append(found, entry.(string))
			}
		}
	}

	actions := []string{"apply", "plan", "modify", "destroy"}

	for _, action := range actions {
		if temp[action] != nil {
			myentries := temp[action].([]interface{})
			for _, entry := range myentries {
				found = append(found, entry.(string))
			}
		}
	}

	return found
}
