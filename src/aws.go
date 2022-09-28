package pike

import (
	"encoding/json"
	"fmt"
)

// GetAWSPermissions for AWS resources
func GetAWSPermissions(result ResourceV2) ([]string, error) {
	var err error
	var Permissions []string
	if result.TypeName == "resource" {
		Permissions, err = GetAWSResourcePermissions(result)
		if err != nil {
			return Permissions, err
		}
	} else {
		Permissions, err = GetAWSDataPermissions(result)
		if err != nil {
			return Permissions, err
		}
	}

	return Permissions, nil
}

// GetAWSResourcePermissions looks up permissions required for resources
func GetAWSResourcePermissions(result ResourceV2) ([]string, error) {

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
		"aws_rds_cluster_instance":                           awsDbInstance,
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
		"aws_default_network_acl":                            awsDefaultNetworkACL,
		"aws_default_route_table":                            awsDefaultRouteTable,
		"aws_default_subnet":                                 awsDefaultSubnet,
		"aws_default_vpc":                                    awsDefaultVpc,
		"aws_default_vpc_dhcp_options":                       awsDefaultVpcDhcpOptions,
		"aws_db_subnet_group":                                awsDbSubnetGroup,
		"aws_docdb_subnet_group":                             awsDbSubnetGroup,
		"aws_neptune_subnet_group":                           awsDbSubnetGroup,
		"aws_wafv2_web_acl":                                  awsWafv2WebACL,
		"aws_wafv2_regex_pattern_set":                        awsWafv2RegexPatternSet,
		"aws_wafv2_rule_group":                               awsWafv2RuleGroup,
		"aws_wafv2_ip_set":                                   awsWafv2IpSet,
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
		"aws_kms_alias":                                      awsKmsAlias,
		"aws_ecr_lifecycle_policy":                           awsEcrLifecyclePolicy,
		"aws_ecr_pull_through_cache_rule":                    awsEcrPullThroughCacheRule,
		"aws_ecr_repository":                                 awsEcrRepository,
		"aws_route_table":                                    awsRouteTable,
		"aws_route_table_association":                        awsRouteTableAssociation,
		"aws_nat_gateway":                                    awsNatGateway,
		"aws_db_option_group":                                awsDbOptionGroup,
		"aws_db_parameter_group":                             awsDbParameterGroup,
		"aws_neptune_parameter_group":                        awsDbParameterGroup,
		"aws_secretsmanager_secret":                          awsSecretsmanagerSecret,
		"aws_secretsmanager_secret_version":                  awsSecretsmanagerSecretVersion,
		"aws_vpc_endpoint":                                   awsVpcEndpoint,
		"aws_vpn_gateway":                                    awsVpnGateway,
		"aws_ssm_document":                                   awsSsmDocument,
		"aws_glue_catalog_database":                          awsGlueCatalogDatabase,
		"aws_glue_catalog_table":                             awsGlueCatalogTable,
		"aws_glue_classifier":                                awsGlueClassifier,
		"aws_glue_crawler":                                   awsGlueCrawler,
		"aws_glue_connection":                                awsGlueConnection,
		"aws_glue_data_catalog_encryption_settings":          awsGlueDataCatalogEncryptionSettings,
		"aws_glue_ml_transform":                              awsGlueMlTransform,
		"aws_glue_trigger":                                   awsGlueTrigger,
		"aws_glue_job":                                       awsGlueJob,
		"aws_glue_registry":                                  awsGlueRegistry,
		"aws_glue_resource_policy":                           awsGlueResourcePolicy,
		"aws_glue_schema":                                    awsGlueSchema,
		"aws_glue_security_configuration":                    awsGlueSecurityConfiguration,
		"aws_glue_user_defined_function":                     awsGlueUserDefinedFunction,
		"aws_glue_workflow":                                  awsGlueWorkflow,
		"aws_codebuild_project":                              awsCodebuildProject,
		"aws_codecommit_repository":                          awsCodecommitRepository,
		"aws_codepipeline":                                   awsCodepipeline,
		"aws_codeartifact_domain":                            awsCodeartifactDomain,
		"aws_codeartifact_domain_permissions_policy":         awsCodeartifactDomainPermissionsPolicy,
		"aws_codeartifact_repository":                        awsCodeartifactRepository,
		"aws_codeartifact_repository_permissions_policy":     awsCodeartifactRepositoryPermissionsPolicy,
		"aws_ssm_patch_baseline":                             awsSsmPatchBaseline,
		"aws_ssm_patch_group":                                awsSsmPatchGroup,
		"aws_ssm_maintenance_window":                         awsSsmMaintenanceWindow,
		"aws_ssm_maintenance_window_target":                  awsSsmMaintenanceWindowTarget,
		"aws_ssm_maintenance_window_task":                    awsSsmMaintenanceWindowTask,
		"aws_launch_template":                                awsLaunchTemplate,
		"aws_directory_service_directory":                    awsDirectoryServiceDirectory,
		"aws_directory_service_log_subscription":             awsDirectoryServiceLogSubscription,
		"aws_cloudtrail":                                     awsCloudtrail,
		"aws_rds_cluster_parameter_group":                    awsRdsClusterParameterGroup,
		"aws_docdb_cluster_parameter_group":                  awsRdsClusterParameterGroup,
		"aws_neptune_cluster_parameter_group":                awsRdsClusterParameterGroup,
		"aws_network_acl_rule":                               awsNetworkACLRule,
		"aws_acm_certificate":                                awsAcmCertificate,
		"aws_acmpca_certificate_authority":                   awsAcmpcaCertificateAuthority,
		"aws_acm_certificate_validation":                     placeholder,
		"aws_ecs_cluster":                                    awsEcsCluster,
		"aws_ecs_service":                                    awsEcsService,
		"aws_ecs_task_definition":                            awsEcsTaskDefinition,
		"aws_appautoscaling_scheduled_action":                awsAppautoscalingScheduledAction,
		"aws_appautoscaling_policy":                          awsAppautoscalingPolicy,
		"aws_appautoscaling_target":                          awsAppautoscalingTarget,
		"aws_cognito_identity_provider":                      awsCognitoIdentityProvider,
		"aws_cognito_resource_server":                        awsCognitoResourceServer,
		"aws_cognito_risk_configuration":                     awsCognitoRiskConfiguration,
		"aws_cognito_user":                                   awsCognitoUser,
		"aws_cognito_user_group":                             awsCognitoUserGroup,
		"aws_cognito_user_in_group":                          awsCognitoUserInGroup,
		"aws_cognito_user_pool":                              awsCognitoUserPool,
		"aws_cognito_user_pool_client":                       awsCognitoUserPoolClient,
		"aws_cognito_user_pool_domain":                       awsCognitoUserPoolDomain,
		"aws_cognito_user_pool_ui_customization":             awsCognitoUserPoolUICustomization,
		"aws_redshift_cluster":                               awsRedshiftCluster,
		"aws_redshift_authentication_profile":                awsRedshiftAuthenticationProfile,
		"aws_redshift_cluster_iam_roles":                     awsRedshiftClusterIamRoles,
		"aws_redshift_event_subscription":                    awsRedshiftEventSubscription,
		"aws_redshift_hsm_client_certificate":                awsRedshiftHsmClientCertififcate,
		"aws_redshift_hsm_configuration":                     awsRedshiftHsmConfiguration,
		"aws_redshift_parameter_group":                       awsRedshiftParameterGroup,
		"aws_redshift_scheduled_action":                      awsRedshiftScheduledAction,
		"aws_redshift_snapshot_copy_grant":                   awsRedshiftSnapshotCopyGrant,
		"aws_redshift_snapshot_schedule":                     awsRedshiftSnapshotSchedule,
		"aws_redshift_snapshot_schedule_association":         awsRedshiftSnapshotScheduleAssociation,
		"aws_redshift_subnet_group":                          awsRedshiftSubnetGroup,
		"aws_redshift_usage_limit":                           awsRedshiftUsageLimit,
		"aws_inspector_assessment_target":                    awsInspectorAssessmentTarget,
		"aws_inspector_assessment_template":                  awsInspectorAssessmentTemplate,
		"aws_inspector_resource_group":                       awsInspectorResouceGroup,
		"aws_efs_access_point":                               awsEfsAccessPoint,
		"aws_efs_backup_policy":                              awsEfsBackupPolicy,
		"aws_efs_file_system":                                awsEfsFileSystem,
		"aws_efs_file_system_policy":                         awsEfsFileSystemPolicy,
		"aws_efs_mount_target":                               awsEfsMountTarget,
		"aws_efs_replication_configuration":                  awsEfsReplicationConfiguration,
		"aws_db_cluster_snapshot":                            awsDbClusterSnapshot,
		"aws_rds_cluster":                                    awsRdsCluster,
		"aws_rds_cluster_activity_stream":                    awsRdsClusterActivityStream,
		"aws_rds_cluster_endpoint":                           awsRdsClusterEndpoint,
		"aws_rds_cluster_role_association":                   awsRdsClusterRoleAssociation,
		"aws_rds_global_cluster":                             awsRdsGlobalCluster,
		"aws_dax_parameter_group":                            awsDaxParameterGroup,
		"aws_dax_subnet_group":                               awsDaxSubnetGroup,
		"aws_elasticache_parameter_group":                    awsElasticacheParameterGroup,
		"aws_elasticache_subnet_group":                       awsElasticacheSubnetGroup,
		"aws_memorydb_subnet_group":                          awsMemorydbSubnetGroup,
		"aws_servicecatalog_portfolio":                       awsServicecatalogPortfolio,
		"aws_sfn_activity":                                   awsSfnActivity,
		"aws_sfn_state_machine":                              awsSfnStateMachine,
		"aws_config_config_rule":                             awsConfigConfigRule,
		"aws_config_configuration_recorder":                  awsConfigConfigurationRecorder,
		"aws_config_configuration_recorder_status":           awsConfigConfigurationRecorderStatus,
		"aws_config_delivery_channel":                        awsConfigDeliveryChannel,
		"aws_ec2_tag":                                        awsEc2Tag,
		"aws_iam_account_alias":                              awsIamAccountAlias,
		"aws_iam_account_password_policy":                    awsIamAccountPasswordPolicy,
		"aws_iam_openid_connect_provider":                    awsIamOpenidConnectProvider,
		"aws_iam_saml_provider":                              awsIamSamlProvider,
		"aws_iam_user_ssh_key":                               awsIamUserSSHKey,
		"aws_elasticache_cluster":                            awsElasticacheCluster,
		"aws_elasticache_replication_group":                  awsElasticacheReplicationGroup,
		"aws_elasticsearch_domain":                           awsElasticsearchDomain,
		"aws_elasticsearch_domain_policy":                    awsElasticsearchDomainPolicy,
		"aws_media_convert_queue":                            awsMediaConvertQueue,
		"aws_imagebuilder_component":                         awsImagebuilderComponent,
		"aws_imagebuilder_distribution_configuration":        awsImagebuilderDistributionConfiguration,
		"aws_imagebuilder_image":                             awsImagebuilderImage,
		"aws_imagebuilder_image_pipeline":                    awsImagebuilderImagePipeline,
		"aws_imagebuilder_image_recipe":                      awsImagebuilderImageRecipe,
		"aws_imagebuilder_infrastructure_configuration":      awsImagebuilderInstrastructureConfiguration,
		"aws_glacier_vault_lock":                             awsGlacierVaultLock,
		"aws_glacier_vault":                                  awsGlacierVault,
		"aws_dlm_lifecycle_policy":                           awsDlmLifecyclePolicy,
		"aws_lambda_layer_version":                           awsLambdaLayerVersion,
		"aws_lambda_layer_version_permission":                awsLambdaLayerVersionPermission,
		"aws_batch_compute_environment":                      awsBatchComputeEnvironment,
		"aws_batch_job_definition":                           awsBatchJobDefinition,
		"aws_batch_job_queue":                                awsBatchJobQueue,
		"aws_batch_scheduling_policy":                        awsBatchSchedulingPolicy,
		"aws_ecr_repository_policy":                          awsEcrRepositoryPolicy,
		"aws_ecrpublic_repository":                           awsEcrPublicRepository,
		"aws_egress_only_internet_gateway":                   awsEgressOnlyInternetGateway,
		"aws_vpc_dhcp_options":                               awsVpcDhcpOptions,
		"aws_vpc_dhcp_options_association":                   awsVpcDhcpOptionsAssociation,
		"aws_vpc_endpoint_service":                           awsVpcEndpointService,
		"aws_vpc_ipv4_cidr_block_association":                awsVpcIpv4CidrBlockAssociation,
		"aws_vpc_peering_connection":                         awsVpcPeeringConnection,
		"aws_vpc_peering_connection_accepter":                awsVpcPeeringConnectionAccepter,
		"aws_vpc_peering_connection_options":                 awsVpcPeeringConnectionOptions,
		"aws_vpn_gateway_attachment":                         awsVpnGatewayAttachment,
		"aws_vpn_gateway_route_propagation":                  awsVpnGatewayRoutePropagation,
		"aws_memorydb_cluster":                               awsMemorydbCluster,
		"aws_memorydb_snapshot":                              awsMemorydbSnapshot,
		"aws_customer_gateway":                               awsCustomerGateway,
		"aws_fsx_openzfs_volume":                             awsFsxOpenzfsVolume,
		"aws_fsx_openzfs_snapshot":                           awsFsxOpenzfsSnaphot,
		"aws_fsx_openzfs_file_system":                        awsFsxOpenzfsFileSystem,
		"aws_codedeploy_app":                                 awsCodedeployApp,
		"aws_codedeploy_deployment_config":                   awsCodedeployDeploymentConfig,
		"aws_codedeploy_deployment_group":                    awsCodedeployDeploymentGroup,
		"aws_networkfirewall_firewall":                       awsNetworkfirewallFirewall,
		"aws_networkfirewall_firewall_policy":                awsNetworkfirewallFirewallPolicy,
		"aws_networkfirewall_logging_configuration":          awsNetworkfirewallLoggingConfiguration,
		"aws_networkfirewall_rule_group":                     awsNetworkfirewallRuleGroup,
	}

	var Permissions []string
	var err error

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("%s not implemented", result.Name)
	}

	return Permissions, err
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
func GetPermissionMap(raw []byte, attributes []string) ([]string, error) {
	var mappings []interface{}
	err := json.Unmarshal(raw, &mappings)
	if err != nil {
		return nil, err
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
			myEntries := temp[action].([]interface{})
			for _, entry := range myEntries {
				found = append(found, entry.(string))
			}
		}
	}

	return found, nil
}
