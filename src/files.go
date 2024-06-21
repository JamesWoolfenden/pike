package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/aws/resource/network-firewall/aws_networkfirewall_resource_policy.json
var awsNetworkfirewallResourcePolicy []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_target_group.json
var awsLbTargetGroup []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_target_group_attachment.json
var awsLbTargetGroupAttachment []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_listener.json
var awsLbListener []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_listener_rule.json
var awsLbListenerRule []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb.json
var awsLb []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket.json
var awsS3Bucket []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_acl.json
var awsS3BucketACL []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_versioning.json
var awsS3BucketVersioning []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_server_side_encryption_configuration.json
var awsS3BucketServerSideEncryptionConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_public_access_block.json
var awsS3BucketPublicAccessBlock []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_logging.json
var awsS3BucketLogging []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_lifecycle_configuration.json
var awsS3BucketLifecycleConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_policy.json
var awsS3BucketPolicy []byte

//go:embed mapping/aws/resource/s3/aws_s3_object.json
var awsS3Object []byte

//go:embed mapping/aws/resource/ec2/aws_instance.json
var awsInstance []byte

//go:embed mapping/aws/resource/ec2/aws_security_group.json
var awsSecurityGroup []byte

//go:embed mapping/aws/resource/ec2/aws_security_group_rule.json
var awsSecurityGroupRule []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_function.json
var awsLambdaFunction []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_alias.json
var awsLambdaAlias []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_permission.json
var awsLambdaPermission []byte

//go:embed mapping/aws/resource/ec2/aws_vpc.json
var awsVpc []byte

//go:embed mapping/aws/resource/ec2/aws_subnet.json
var awsSubnet []byte

//go:embed mapping/aws/resource/ec2/aws_network_acl.json
var awsNetworkACL []byte

//go:embed mapping/aws/resource/kms/aws_kms_key.json
var awsKmsKey []byte

//go:embed mapping/aws/resource/iam/aws_iam_role.json
var awsIamRole []byte

//go:embed mapping/aws/resource/iam/aws_iam_role_policy.json
var awsIamRolePolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_role_policy_attachment.json
var awsIamRolePolicyAttachment []byte

//go:embed mapping/aws/resource/iam/aws_iam_policy.json
var awsIamPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_instance_profile.json
var awsIamInstanceProfile []byte

//go:embed mapping/aws/resource/iam/aws_iam_access_key.json
var awsIamAccessKey []byte

//go:embed mapping/aws/resource/iam/aws_iam_group.json
var awsIamGroup []byte

//go:embed mapping/aws/resource/iam/aws_iam_group_membership.json
var awsIamGroupMembership []byte

//go:embed mapping/aws/resource/iam/aws_iam_group_policy.json
var awsIamGroupPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_group_policy_attachment.json
var awsIamGroupPolicyAttachment []byte

//go:embed mapping/aws/resource/iam/aws_iam_policy_attachment.json
var awsIamPolicyAttachment []byte

//go:embed mapping/aws/resource/iam/aws_iam_service_linked_role.json
var awsIamServiceLinkedRole []byte

//go:embed mapping/aws/resource/iam/aws_iam_server_certificate.json
var awsIamServerCertificate []byte

//go:embed mapping/aws/resource/iam/aws_iam_user.json
var awsIamUser []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_login_profile.json
var awsIamUserLoginProfile []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_policy.json
var awsIamUserPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_policy_attachment.json
var awsIamUserPolicyAttachment []byte

//go:embed mapping/aws/resource/mq/aws_mq_broker.json
var awsMqBroker []byte

//go:embed mapping/aws/resource/mq/aws_mq_configuration.json
var awsMqConfiguration []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_group.json
var awsCloudwatchLogGroup []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_stream.json
var awsCloudwatchLogStream []byte

//go:embed mapping/aws/resource/events/aws_cloudwatch_event_rule.json
var awsCloudwatchEventRule []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_metric_filter.json
var awsCloudwatchLogMetricFilter []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_resource_policy.json
var awsCloudwatchLogResourcePolicy []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_subscription_filter.json
var awsCloudwatchLogSubscriptionFilter []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_metric_alarm.json
var awsCloudwatchMetricAlarm []byte

//go:embed mapping/aws/resource/events/aws_cloudwatch_event_target.json
var awsCloudwatchEventTarget []byte

//go:embed mapping/aws/resource/route53/aws_route53_record.json
var awsRoute53Record []byte

//go:embed mapping/aws/resource/route53/aws_route53_zone.json
var awsRoute53Zone []byte

//go:embed mapping/aws/resource/route53/aws_route53_zone_association.json
var awsRoute53ZoneAssociation []byte

//go:embed mapping/aws/resource/sns/aws_sns_topic.json
var awsSnsTopic []byte

//go:embed mapping/aws/resource/sns/aws_sns_topic_subscription.json
var awsSnsTopicSubscription []byte

//go:embed mapping/aws/resource/sns/aws_sns_topic_policy.json
var awsSnsTopicPolicy []byte

//go:embed mapping/aws/resource/ec2/aws_key_pair.json
var awsKeyPair []byte

//go:embed mapping/aws/resource/rds/aws_db_instance.json
var awsDBInstance []byte

//go:embed mapping/aws/resource/rds/aws_db_cluster_snapshot.json
var awsDBClusterSnapshot []byte

//go:embed mapping/aws/resource/rds/aws_db_event_subscription.json
var awsDBEventSubscription []byte

//go:embed mapping/aws/resource/ram/aws_ram_principal_association.json
var awsRAMPrincipleAssociation []byte

//go:embed mapping/aws/resource/ram/aws_ram_resource_association.json
var awsRAMResourceAssociation []byte

//go:embed mapping/aws/resource/ram/aws_ram_resource_share.json
var awsRAMResourceShare []byte

//go:embed mapping/aws/resource/rds/aws_rds_cluster.json
var awsRdsCluster []byte

//go:embed mapping/aws/resource/rds/aws_rds_cluster_activity_stream.json
var awsRdsClusterActivityStream []byte

//go:embed mapping/aws/resource/rds/aws_rds_cluster_endpoint.json
var awsRdsClusterEndpoint []byte

//go:embed mapping/aws/resource/rds/aws_rds_cluster_role_association.json
var awsRdsClusterRoleAssociation []byte

//go:embed mapping/aws/resource/rds/aws_rds_global_cluster.json
var awsRdsGlobalCluster []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_table.json
var awsDynamodbTable []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_parameter.json
var awsSsmParameter []byte

//go:embed mapping/aws/resource/ec2/aws_route.json
var awsRoute []byte

//go:embed mapping/aws/resource/ec2/aws_default_security_group.json
var awsDefaultSecurityGroup []byte

//go:embed mapping/aws/resource/ec2/aws_default_network_acl.json
var awsDefaultNetworkACL []byte

//go:embed mapping/aws/resource/ec2/aws_default_subnet.json
var awsDefaultSubnet []byte

//go:embed mapping/aws/resource/ec2/aws_default_vpc.json
var awsDefaultVpc []byte

//go:embed mapping/aws/resource/ec2/aws_default_route_table.json
var awsDefaultRouteTable []byte

//go:embed mapping/aws/resource/ec2/aws_default_vpc_dhcp_options.json
var awsDefaultVpcDhcpOptions []byte

//go:embed mapping/aws/resource/rds/aws_db_subnet_group.json
var awsDBSubnetGroup []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_web_acl.json
var awsWafv2WebACL []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_ip_set.json
var awsWafv2IpSet []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_rule_group.json
var awsWafv2RuleGroup []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_regex_pattern_set.json
var awsWafv2RegexPatternSet []byte

//go:embed mapping/aws/resource/ec2/aws_vpn_connection.json
var awsVpnConnection []byte

//go:embed mapping/aws/resource/ec2/aws_vpn_gateway.json
var awsVpnGateway []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint.json
var awsVpcEndpoint []byte

//go:embed mapping/aws/resource/apigateway/aws_api_gateway_rest_api.json
var awsAPIGatewayRestAPI []byte

//go:embed mapping/aws/resource/apigateway/aws_apigatewayv2_api.json
var awsApigatewayv2Api []byte

//go:embed mapping/aws/resource/apigateway/aws_api_gateway_account.json
var awsAPIGatewayAccount []byte

//go:embed mapping/aws/resource/sqs/aws_sqs_queue.json
var awsSqsQueue []byte

//go:embed mapping/aws/resource/sqs/aws_sqs_queue_policy.json
var awsSqsQueuePolicy []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_volume.json
var awsEbsVolume []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_group.json
var awsAutoscalingGroup []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_lifecycle_hook.json
var awsAutoscalingLifecycleHook []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_notification.json
var awsAutoscalingNotification []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_policy.json
var awsAutoscalingPolicy []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_attachment.json
var awsAutoscalingAttachment []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_elb.json
var awsElb []byte

//go:embed mapping/aws/resource/ec2/aws_internet_gateway.json
var awsInternetGateway []byte

//go:embed mapping/aws/resource/autoscaling/aws_launch_configuration.json
var awsLaunchConfiguration []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_capacity_reservation.json
var awsEc2CapacityReservation []byte

//go:embed mapping/aws/resource/ec2/aws_network_interface.json
var awsNetworkInterface []byte

//go:embed mapping/aws/resource/ec2/aws_placement_group.json
var awsPlacementGroup []byte

//go:embed mapping/aws/resource/ec2/aws_spot_instance_request.json
var awsSpotInstanceRequest []byte

//go:embed mapping/aws/resource/ec2/aws_volume_attachment.json
var awsVolumeAttachment []byte

//go:embed mapping/aws/resource/ec2/aws_eip.json
var awsEip []byte

//go:embed mapping/aws/resource/firehose/aws_kinesis_firehose_delivery_stream.json
var awsKinesisFirehoseDeliveryStream []byte

//go:embed mapping/aws/resource/firehose/aws_kinesis_stream.json
var awsKinesisStream []byte

//go:embed mapping/aws/resource/firehose/aws_kinesis_video_stream.json
var awsKinesisVideoStream []byte

//go:embed mapping/aws/resource/budgets/aws_budgets_budget.json
var awsBudgetsBudget []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_elastic_beanstalk_application.json
var awsElasticBeanstalkApplication []byte

//go:embed mapping/aws/resource/ec2/aws_flow_log.json
var awsFlowLog []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_lifecycle_policy.json
var awsEcrLifecyclePolicy []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_pull_through_cache_rule.json
var awsEcrPullThroughCacheRule []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_registry_scanning_configuration.json
var awsEcrRegistryScanningConfiguration []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_repository.json
var awsEcrRepository []byte

//go:embed mapping/aws/resource/kms/aws_kms_alias.json
var awsKmsAlias []byte

//go:embed mapping/aws/resource/ec2/aws_route_table.json
var awsRouteTable []byte

//go:embed mapping/aws/resource/ec2/aws_route_table_association.json
var awsRouteTableAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_nat_gateway.json
var awsNatGateway []byte

//go:embed mapping/aws/resource/rds/aws_db_option_group.json
var awsDBOptionGroup []byte

//go:embed mapping/aws/resource/rds/aws_db_parameter_group.json
var awsDBParameterGroup []byte

//go:embed mapping/aws/resource/secretsmanager/aws_secretsmanager_secret.json
var awsSecretsmanagerSecret []byte

//go:embed mapping/aws/resource/secretsmanager/aws_secretsmanager_secret_version.json
var awsSecretsmanagerSecretVersion []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_document.json
var awsSsmDocument []byte

//go:embed mapping/aws/resource/glue/aws_glue_classifier.json
var awsGlueClassifier []byte

//go:embed mapping/aws/resource/glue/aws_glue_crawler.json
var awsGlueCrawler []byte

//go:embed mapping/aws/resource/glue/aws_glue_catalog_database.json
var awsGlueCatalogDatabase []byte

//go:embed mapping/aws/resource/glue/aws_glue_catalog_table.json
var awsGlueCatalogTable []byte

//go:embed mapping/aws/resource/glue/aws_glue_connection.json
var awsGlueConnection []byte

//go:embed mapping/aws/resource/glue/aws_glue_data_catalog_encryption_settings.json
var awsGlueDataCatalogEncryptionSettings []byte

//go:embed mapping/aws/resource/glue/aws_glue_ml_transform.json
var awsGlueMlTransform []byte

//go:embed mapping/aws/resource/glue/aws_glue_trigger.json
var awsGlueTrigger []byte

//go:embed mapping/aws/resource/glue/aws_glue_job.json
var awsGlueJob []byte

//go:embed mapping/aws/resource/glue/aws_glue_registry.json
var awsGlueRegistry []byte

//go:embed mapping/aws/resource/glue/aws_glue_resource_policy.json
var awsGlueResourcePolicy []byte

//go:embed mapping/aws/resource/glue/aws_glue_schema.json
var awsGlueSchema []byte

//go:embed mapping/aws/resource/glue/aws_glue_security_configuration.json
var awsGlueSecurityConfiguration []byte

//go:embed mapping/aws/resource/glue/aws_glue_user_defined_function.json
var awsGlueUserDefinedFunction []byte

//go:embed mapping/aws/resource/glue/aws_glue_workflow.json
var awsGlueWorkflow []byte

//go:embed mapping/aws/resource/codebuild/aws_codebuild_project.json
var awsCodebuildProject []byte

//go:embed mapping/aws/resource/codecommit/aws_codecommit_repository.json
var awsCodecommitRepository []byte

//go:embed mapping/aws/resource/codepipeline/aws_codepipeline.json
var awsCodepipeline []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_repository.json
var awsCodeartifactRepository []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_domain.json
var awsCodeartifactDomain []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_repository_permissions_policy.json
var awsCodeartifactRepositoryPermissionsPolicy []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_domain_permissions_policy.json
var awsCodeartifactDomainPermissionsPolicy []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_patch_baseline.json
var awsSsmPatchBaseline []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_patch_group.json
var awsSsmPatchGroup []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_maintenance_window.json
var awsSsmMaintenanceWindow []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_maintenance_window_target.json
var awsSsmMaintenanceWindowTarget []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_maintenance_window_task.json
var awsSsmMaintenanceWindowTask []byte

//go:embed mapping/aws/resource/ec2/aws_launch_template.json
var awsLaunchTemplate []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_directory.json
var awsDirectoryServiceDirectory []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_log_subscription.json
var awsDirectoryServiceLogSubscription []byte

//go:embed mapping/aws/resource/cloudtrail/aws_cloudtrail.json
var awsCloudtrail []byte

//go:embed mapping/aws/resource/rds/aws_rds_cluster_parameter_group.json
var awsRdsClusterParameterGroup []byte

//go:embed mapping/aws/resource/ec2/aws_network_acl_rule.json
var awsNetworkACLRule []byte

//go:embed mapping/aws/resource/acm/aws_acm_certificate.json
var AWSAcmCertificate []byte

//go:embed mapping/aws/resource/acm-pa/aws_acmpca_certificate_authority.json
var awsAcmpcaCertificateAuthority []byte

//go:embed mapping/aws/resource/acm-pa/aws_acmpca_certificate_authority_certificate.json
var awsAcmpcaCertificateAuthorityCertificate []byte

//go:embed mapping/aws/resource/acm-pa/aws_acmpca_certificate.json
var awsAcmpcaCertificate []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_cluster.json
var awsEcsCluster []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_task_definition.json
var awsEcsTaskDefinition []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_service.json
var awsEcsService []byte

//go:embed mapping/aws/resource/application-autoscaling/aws_appautoscaling_scheduled_action.json
var awsAppautoscalingScheduledAction []byte

//go:embed mapping/aws/resource/application-autoscaling/aws_appautoscaling_target.json
var awsAppautoscalingTarget []byte

//go:embed mapping/aws/resource/application-autoscaling/aws_appautoscaling_policy.json
var awsAppautoscalingPolicy []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_identity_provider.json
var awsCognitoIdentityProvider []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_resource_server.json
var awsCognitoResourceServer []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_identity_provider.json
var awsCognitoRiskConfiguration []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user.json
var awsCognitoUser []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_group.json
var awsCognitoUserGroup []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_in_group.json
var awsCognitoUserInGroup []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool.json
var awsCognitoUserPool []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool_client.json
var awsCognitoUserPoolClient []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool_domain.json
var awsCognitoUserPoolDomain []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool_ui_customization.json
var awsCognitoUserPoolUICustomization []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_subnet_group.json
var awsRedshiftSubnetGroup []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_cluster.json
var awsRedshiftCluster []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_scheduled_action.json
var awsRedshiftScheduledAction []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_authentication_profile.json
var awsRedshiftAuthenticationProfile []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_cluster_iam_roles.json
var awsRedshiftClusterIamRoles []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_event_subscription.json
var awsRedshiftEventSubscription []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_hsm_client_certificate.json
var awsRedshiftHsmClientCertififcate []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_hsm_configuration.json
var awsRedshiftHsmConfiguration []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_usage_limit.json
var awsRedshiftUsageLimit []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_parameter_group.json
var awsRedshiftParameterGroup []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_snapshot_copy_grant.json
var awsRedshiftSnapshotCopyGrant []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_snapshot_schedule.json
var awsRedshiftSnapshotSchedule []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_snapshot_schedule_association.json
var awsRedshiftSnapshotScheduleAssociation []byte

//go:embed mapping/aws/resource/inspector/aws_inspector_assessment_target.json
var awsInspectorAssessmentTarget []byte

//go:embed mapping/aws/resource/inspector/aws_inspector_assessment_template.json
var awsInspectorAssessmentTemplate []byte

//go:embed mapping/aws/resource/inspector/aws_inspector_resource_group.json
var awsInspectorResouceGroup []byte

//go:embed mapping/aws/resource/elasticfilesystem/aws_efs_access_point.json
var awsEfsAccessPoint []byte

//go:embed mapping/aws/resource/elasticfilesystem/aws_efs_file_system.json
var awsEfsFileSystem []byte

//go:embed mapping/aws/resource/elasticfilesystem/aws_efs_backup_policy.json
var awsEfsBackupPolicy []byte

//go:embed mapping/aws/resource/elasticfilesystem/aws_efs_file_system_policy.json
var awsEfsFileSystemPolicy []byte

//go:embed mapping/aws/resource/elasticfilesystem/aws_efs_mount_target.json
var awsEfsMountTarget []byte

//go:embed mapping/aws/resource/elasticfilesystem/aws_efs_replication_configuration.json
var awsEfsReplicationConfiguration []byte

//go:embed mapping/aws/resource/dax/aws_dax_parameter_group.json
var awsDaxParameterGroup []byte

//go:embed mapping/aws/resource/dax/aws_dax_subnet_group.json
var awsDaxSubnetGroup []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_parameter_group.json
var awsElasticacheParameterGroup []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_subnet_group.json
var awsElasticacheSubnetGroup []byte

//go:embed mapping/aws/resource/memorydb/aws_memorydb_subnet_group.json
var awsMemorydbSubnetGroup []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_portfolio.json
var awsServicecatalogPortfolio []byte

//go:embed mapping/aws/resource/states/aws_sfn_activity.json
var awsSfnActivity []byte

//go:embed mapping/aws/resource/states/aws_sfn_state_machine.json
var awsSfnStateMachine []byte

//go:embed mapping/aws/resource/config/aws_config_config_rule.json
var awsConfigConfigRule []byte

//go:embed mapping/aws/resource/config/aws_config_configuration_aggregator.json
var awsConfigConfigurationAggregator []byte

//go:embed mapping/aws/resource/config/aws_config_configuration_recorder_status.json
var awsConfigConfigurationRecorderStatus []byte

//go:embed mapping/aws/resource/config/aws_config_configuration_recorder.json
var awsConfigConfigurationRecorder []byte

//go:embed mapping/aws/resource/config/aws_config_delivery_channel.json
var awsConfigDeliveryChannel []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_tag.json
var awsEc2Tag []byte

//go:embed mapping/aws/resource/iam/aws_iam_account_alias.json
var awsIamAccountAlias []byte

//go:embed mapping/aws/resource/iam/aws_iam_account_password_policy.json
var awsIamAccountPasswordPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_openid_connect_provider.json
var awsIamOpenidConnectProvider []byte

//go:embed mapping/aws/resource/iam/aws_iam_saml_provider.json
var awsIamSamlProvider []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_ssh_key.json
var awsIamUserSSHKey []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_cluster.json
var awsElasticacheCluster []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_replication_group.json
var awsElasticacheReplicationGroup []byte

//go:embed mapping/aws/resource/es/aws_elasticsearch_domain.json
var awsElasticsearchDomain []byte

//go:embed mapping/aws/resource/es/aws_elasticsearch_domain_policy.json
var awsElasticsearchDomainPolicy []byte

//go:embed mapping/aws/resource/mediaconvert/aws_media_convert_queue.json
var awsMediaConvertQueue []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_component.json
var awsImagebuilderComponent []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_distribution_configuration.json
var awsImagebuilderDistributionConfiguration []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_image.json
var awsImagebuilderImage []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_image_pipeline.json
var awsImagebuilderImagePipeline []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_image_recipe.json
var awsImagebuilderImageRecipe []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_infrastructure_configuration.json
var awsImagebuilderInstrastructureConfiguration []byte

//go:embed mapping/aws/resource/glacier/aws_glacier_vault_lock.json
var awsGlacierVaultLock []byte

//go:embed mapping/aws/resource/glacier/aws_glacier_vault.json
var awsGlacierVault []byte

//go:embed mapping/aws/resource/dlm/aws_dlm_lifecycle_policy.json
var awsDlmLifecyclePolicy []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_layer_version.json
var awsLambdaLayerVersion []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_layer_version_permission.json
var awsLambdaLayerVersionPermission []byte

//go:embed mapping/aws/resource/batch/aws_batch_compute_environment.json
var awsBatchComputeEnvironment []byte

//go:embed mapping/aws/resource/batch/aws_batch_job_definition.json
var awsBatchJobDefinition []byte

//go:embed mapping/aws/resource/batch/aws_batch_job_queue.json
var awsBatchJobQueue []byte

//go:embed mapping/aws/resource/batch/aws_batch_scheduling_policy.json
var awsBatchSchedulingPolicy []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_repository_policy.json
var awsEcrRepositoryPolicy []byte

//go:embed mapping/aws/resource/ecr-public/aws_ecrpublic_repository.json
var awsEcrPublicRepository []byte

//go:embed mapping/aws/resource/ec2/aws_egress_only_internet_gateway.json
var awsEgressOnlyInternetGateway []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_dhcp_options.json
var awsVpcDhcpOptions []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_dhcp_options_association.json
var awsVpcDhcpOptionsAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_service.json
var awsVpcEndpointService []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_subnet_association.json
var awsVpcEndpointSubnetAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipv4_cidr_block_association.json
var awsVpcIpv4CidrBlockAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_peering_connection.json
var awsVpcPeeringConnection []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_peering_connection_accepter.json
var awsVpcPeeringConnectionAccepter []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_peering_connection_options.json
var awsVpcPeeringConnectionOptions []byte

//go:embed mapping/aws/resource/ec2/aws_vpn_gateway_attachment.json
var awsVpnGatewayAttachment []byte

//go:embed mapping/aws/resource/ec2/aws_vpn_gateway_route_propagation.json
var awsVpnGatewayRoutePropagation []byte

//go:embed mapping/aws/resource/memorydb/aws_memorydb_cluster.json
var awsMemorydbCluster []byte

//go:embed mapping/aws/resource/memorydb/aws_memorydb_snapshot.json
var awsMemorydbSnapshot []byte

//go:embed mapping/aws/resource/ec2/aws_customer_gateway.json
var awsCustomerGateway []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_openzfs_file_system.json
var awsFsxOpenzfsFileSystem []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_openzfs_volume.json
var awsFsxOpenzfsVolume []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_openzfs_snapshot.json
var awsFsxOpenzfsSnaphot []byte

//go:embed mapping/aws/resource/codedeploy/aws_codedeploy_app.json
var awsCodedeployApp []byte

//go:embed mapping/aws/resource/codedeploy/aws_codedeploy_deployment_config.json
var awsCodedeployDeploymentConfig []byte

//go:embed mapping/aws/resource/codedeploy/aws_codedeploy_deployment_group.json
var awsCodedeployDeploymentGroup []byte

//go:embed mapping/aws/resource/network-firewall/aws_networkfirewall_firewall.json
var awsNetworkfirewallFirewall []byte

//go:embed mapping/aws/resource/network-firewall/aws_networkfirewall_firewall_policy.json
var awsNetworkfirewallFirewallPolicy []byte

//go:embed mapping/aws/resource/network-firewall/aws_networkfirewall_logging_configuration.json
var awsNetworkfirewallLoggingConfiguration []byte

//go:embed mapping/aws/resource/network-firewall/aws_networkfirewall_rule_group.json
var awsNetworkfirewallRuleGroup []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_instance.json
var awsLightsailInstance []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_instance_public_ports.json
var awsLightsailInstancePublicPorts []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_key_pair.json
var awsLightsailKeyPair []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_static_ip.json
var awsLightsailStaticIP []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_static_ip_attachment.json
var awsLightsailStaticIPAttachment []byte

//go:embed mapping/aws/resource/medialive/aws_medialive_input.json
var awsMedialiveInput []byte

//go:embed mapping/aws/resource/medialive/aws_medialive_input_security_group.json
var awsMedialiveInputSecurityGroup []byte

//go:embed mapping/aws/resource/cloudformation/aws_cloudformation_stack_set.json
var awsCloudFormationStackSet []byte

//go:embed mapping/aws/resource/cloudformation/aws_cloudformation_stack_set_instance.json
var awsCloudFormationStackSetInstance []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_origin_access_control.json
var awsCloudfrontOriginAccessControl []byte

//go:embed mapping/aws/resource/kafka/aws_msk_serverless_cluster.json
var awsMskServerlessCluster []byte

//go:embed mapping/aws/resource/kafka/aws_msk_cluster.json
var awsMskCluster []byte

//go:embed mapping/aws/resource/kafka/aws_msk_configuration.json
var awsMskConfiguration []byte

//go:embed mapping/aws/resource/kafka/aws_msk_scram_secret_association.json
var awsMskScramSecretAssociation []byte

//go:embed mapping/aws/resource/route53/aws_route53_delegation_set.json
var awsRoute53DelegationSet []byte

//go:embed mapping/aws/resource/route53/aws_route53_key_signing_key.json
var awsRoute53KeySiginingKey []byte

//go:embed mapping/aws/resource/route53/aws_route53_health_check.json
var awsRoute53HealthCheck []byte

//go:embed mapping/aws/resource/route53/aws_route53_hosted_zone_dnssec.json
var awsRoute53HostedZoneDnssec []byte

//go:embed mapping/aws/resource/route53/aws_route53_query_log.json
var awsRoute53QueryLog []byte

//go:embed mapping/aws/resource/route53/aws_route53_vpc_association_authorization.json
var awsRoute53VpcAssociationAuthorization []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_query_log_config.json
var awsRoute53ResolverQueryLogConfig []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_query_log_config_association.json
var awsRoute53ResolverQueryLogConfigAssociation []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_rule_association.json
var awsRoute53ResolverRuleAssociation []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_endpoint_configuration.json
var awsSagemakerEndpointConfiguration []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_model.json
var awsSagemakerModel []byte

//go:embed mapping/aws/resource/sqs/aws_sqs_queue_redrive_allow_policy.json
var awsSqsQueueRedriveAllowPolicy []byte

//go:embed mapping/aws/resource/sqs/aws_sqs_queue_redrive_policy.json
var awsSqsQueueRedrivePolicy []byte

//go:embed mapping/aws/resource/grafana/aws_grafana_workspace_api_key.json
var awsGrafanaWorkspaceAPIKey []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_route.json
var awsEc2TransitGatewayRoute []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway.json
var awsEc2TransitGateway []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_route_table.json
var awsEc2TransitGatewayRouteTable []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_route_table_association.json
var awsEc2TransitGatewayRouteTableAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_route_table_propagation.json
var awsEc2TransitGatewayRouteTablePropagation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_vpc_attachment.json
var awsEc2TransitGatewayVpcAttachment []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_route_table_association.json
var awsVpcEndpointRouteTableAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_network_insights_path.json
var awsEc2NetworkInsightsPath []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_network_insights_analysis.json
var awsEc2NetworkInsightsAnalysis []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_configuration_profile.json
var awsAppconfigConfigurationProfile []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_application.json
var awsAppconfigApplication []byte

//go:embed mapping/aws/resource/dax/aws_dax_cluster.json
var awsDaxCluster []byte

//go:embed mapping/aws/resource/eks/aws_eks_cluster.json
var awsEksCluster []byte

//go:embed mapping/aws/resource/eks/aws_eks_addon.json
var awsEksAddon []byte

//go:embed mapping/aws/resource/eks/aws_eks_node_group.json
var awsEksNodeGroup []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_event_source_mapping.json
var awsLambdaEventSourceMapping []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_function_event_invoke_config.json
var awsLambdaFunctionEventInvokeConfig []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_function_url.json
var awsLambdaFunctionURL []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_provisioned_concurrency_config.json
var awsLambdaProvisionedConcurrencyConfig []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_accelerate_configuration.json
var awsS3BucketAccelerateConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_cors_configuration.json
var awsS3BucketCorsConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_intelligent_tiering_configuration.json
var awsS3BucketIntelligentTieringConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_metric.json
var awsS3BucketMetric []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_object_lock_configuration.json
var awsS3BucketObjectLockCOnfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_ownership_controls.json
var awsS3BucketOwnershipControls []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_replication_configuration.json
var awsS3BucketReplicationConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_request_payment_configuration.json
var awsS3BucketRequestPaymentConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_website_configuration.json
var awsS3BucketWebsiteConfiguration []byte

//go:embed mapping/aws/resource/cloud9/aws_cloud9_environment_ec2.json
var awsCloud9EnvironmentEc2 []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_notification.json
var awsS3BucketNotification []byte

//go:embed mapping/aws/resource/rds/aws_neptune_cluster.json
var awsNeptuneCluster []byte

//go:embed mapping/aws/resource/rds/aws_neptune_cluster_endpoint.json
var awsNeptuneClusterEndpoint []byte

//go:embed mapping/aws/resource/rds/aws_neptune_cluster_instance.json
var awsNeptuneClusterInstance []byte

//go:embed mapping/aws/resource/rds/aws_neptune_event_subscription.json
var awsNeptuneEventSubscription []byte

//go:embed mapping/aws/resource/rds/aws_neptune_cluster_snapshot.json
var awsNeptuneClusterSnapshot []byte

//go:embed mapping/aws/resource/athena/aws_athena_data_catalog.json
var awsAthenaDataCatalog []byte

//go:embed mapping/aws/resource/athena/aws_athena_database.json
var awsAthenaDatabase []byte

//go:embed mapping/aws/resource/athena/aws_athena_workgroup.json
var awsAthenaWorkgroup []byte

//go:embed mapping/aws/resource/athena/aws_athena_named_query.json
var awsAthenaNamedQuery []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_composite_alarm.json
var awsCloudwatchCompositeAlarm []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_dashboard.json
var awsCloudwatchDashboard []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_event_api_destination.json
var awsCloudwatchEventAPIDestination []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_event_archive.json
var awsCloudwatchEventArchive []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_event_bus.json
var awsCloudwatchEventBus []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_event_connection.json
var awsCloudwatchEventConnection []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_event_permission.json
var awsCloudwatchEventPermission []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_distribution.json
var awsCloudfrontDistribution []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_field_level_encryption_config.json
var awsCloudfrontFieldLevelEncryptionConfig []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_field_level_encryption_profile.json
var awsCloudfrontFieldLevelEncryptionProfile []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_key_group.json
var awsCloudfrontKeyGroup []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_monitoring_subscription.json
var awsCloudfrontMonitoringSubscription []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_public_key.json
var awsCloudfrontPublicKey []byte

//go:embed mapping/aws/resource/backup/aws_backup_framework.json
var awsBackupFramework []byte

//go:embed mapping/aws/resource/backup/aws_backup_global_settings.json
var awsBackupGlobalSettings []byte

//go:embed mapping/aws/resource/backup/aws_backup_plan.json
var awsBackupPlan []byte

//go:embed mapping/aws/resource/backup/aws_backup_region_settings.json
var awsBackupRegionSettings []byte

//go:embed mapping/aws/resource/backup/aws_backup_report_plan.json
var awsBackupReportPlan []byte

//go:embed mapping/aws/resource/backup/aws_backup_selection.json
var awsBackupSelection []byte

//go:embed mapping/aws/resource/backup/aws_backup_vault.json
var awsBackupVault []byte

//go:embed mapping/aws/resource/backup/aws_backup_vault_lock_configuration.json
var awsBackupVaultLockConfiguration []byte

//go:embed mapping/aws/resource/backup/aws_backup_vault_notifications.json
var awsBackupVaultNotification []byte

//go:embed mapping/aws/resource/backup/aws_backup_vault_policy.json
var awsBackupVaultPolicy []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_snapshot.json
var awsEbsSnapshot []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_snapshot_copy.json
var awsEbsSnapshotCopy []byte

//go:embed mapping/aws/resource/xray/aws_xray_encryption_config.json
var awsXrayEncryptionConfig []byte

//go:embed mapping/aws/resource/xray/aws_xray_group.json
var awsXrayGroup []byte

//go:embed mapping/aws/resource/xray/aws_xray_sampling_rule.json
var awsXraySamplingRule []byte

//go:embed mapping/aws/resource/kms/aws_kms_grant.json
var awsKmsGrant []byte

//go:embed mapping/aws/resource/applicationinsights/aws_applicationinsights_application.json
var awsApplicationinsightsApplication []byte

//go:embed mapping/aws/resource/resource-groups/aws_resourcegroups_group.json
var awsResourcegroupsGroup []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_inventory.json
var awsS3BucketInventory []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_group_membership.json
var awsIamUserGroupMembership []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_invocation.json
var awsLambdaInvocation []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_user_group.json
var awsElasticacheUserGroup []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_user.json
var awsElasticacheUser []byte

//go:embed mapping/aws/resource/servicequotas/aws_servicequotas_service_quota.json
var awsServicequotasServiceQuota []byte

//go:embed mapping/aws/resource/ses/aws_ses_domain_dkim.json
var awsSesDomainDkim []byte

//go:embed mapping/aws/resource/ses/aws_ses_domain_identity.json
var awsSesDomainIdentity []byte

//go:embed mapping/aws/resource/ses/aws_ses_domain_identity_verification.json
var awsSesDomainIdentityVerification []byte

//go:embed mapping/aws/resource/ses/aws_ses_domain_mail_from.json
var awsSesDomainMailFrom []byte

//go:embed mapping/aws/resource/ses/aws_ses_identity_notification_topic.json
var awsSesIdentityNotificationTopic []byte

//go:embed mapping/aws/resource/ses/aws_ses_identity_policy.json
var awsSesIdentityPolicy []byte

//go:embed mapping/aws/resource/ses/aws_ses_receipt_rule.json
var awsSesReceiptRule []byte

//go:embed mapping/aws/resource/ses/aws_ses_receipt_rule_set.json
var awsSesReceiptRuleSet []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_origin_access_identity.json
var awsCloudfrontOriginAccessIdentity []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_metric_stream.json
var awsCloudwatchMetricStream []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_hosted_transit_virtual_interface_accepter.json
var awsDxHostedTransitVirtualInterfaceAccepter []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_gateway.json
var awsDxGateway []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_gateway_association.json
var awsDxGatewayAssociation []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_global_table.json
var awsDynamodbGlobalTable []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_contributor_insights.json
var awsDynamodbContributorInsights []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_table_item.json
var awsDynamodbTableItem []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_tag.json
var awsDynamodbTag []byte

//go:embed mapping/aws/resource/backend/s3.json
var s3backend []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_response_headers_policy.json
var awsCloudfrontResponseHeadersPolicy []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_cluster.json
var awsEmrCluster []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_security_configuration.json
var awsEmrSecurityConfiguration []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_web_acl_association.json
var awsWafv2WebACLAssociation []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_web_acl_logging_configuration.json
var awsWafv2WebACLLoggingConfiguration []byte

//go:embed mapping/aws/resource/workspaces/aws_workspaces_workspace.json
var awsWorkspacesWorkspace []byte

//go:embed mapping/aws/resource/workspaces/aws_workspaces_directory.json
var awsWorkspacesDirectory []byte

//go:embed mapping/aws/resource/account/aws_account_alternate_contact.json
var awsAccountAlternativeContact []byte

//go:embed mapping/aws/resource/account/aws_account_primary_contact.json
var awsAccountPrimaryContact []byte

//go:embed mapping/aws/resource/sns/aws_sns_sms_preferences.json
var awsSnsSmsPreferences []byte

//go:embed mapping/aws/resource/sns/aws_sns_topic_data_protection_policy.json
var awsSnsTopicDataProtection []byte

//go:embed mapping/aws/resource/ssm-contacts/aws_ssmcontacts_contact.json
var awsSsmcontactsContact []byte

//go:embed mapping/aws/resource/ssm-contacts/aws_ssmcontacts_contact_channel.json
var awsSsmcontactContactChannel []byte

//go:embed mapping/aws/resource/ssm-contacts/aws_ssmcontacts_plan.json
var awsSsmcontactsPlan []byte

//go:embed mapping/aws/resource/ssm-incidents/aws_ssmincidents_replication_set.json
var awsSsmincidentsReplicationSet []byte

//go:embed mapping/aws/resource/swf/aws_swf_domain.json
var awsSwfDomain []byte

//go:embed mapping/aws/resource/ec2/aws_ami.json
var awsAmi []byte

//go:embed mapping/aws/resource/ec2/aws_ami_copy.json
var awsAmiCopy []byte

//go:embed mapping/aws/resource/ec2/aws_ami_from_instance.json
var awsAmiFromInstance []byte

//go:embed mapping/aws/resource/ec2/aws_ami_launch_permission.json
var awsAmiLauchPermission []byte

//go:embed mapping/aws/resource/budgets/aws_budgets_budget_action.json
var awsBudgetsBudgetAction []byte

//go:embed mapping/aws/resource/cloudformation/aws_cloudformation_stack.json
var awsCloudformationStack []byte

//go:embed mapping/aws/resource/cloudformation/aws_cloudformation_type.json
var awsCloudformationType []byte

//go:embed mapping/aws/resource/organizations/aws_organizations_policy.json
var awsOrganizationsPolicy []byte

//go:embed mapping/aws/resource/organizations/aws_organizations_policy_attachment.json
var awsOrganizationsPolicyAttachment []byte

//go:embed mapping/aws/resource/secretsmanager/aws_secretsmanager_secret_policy.json
var awsSecretsmanagerSecretPolicy []byte

//go:embed mapping/aws/resource/secretsmanager/aws_secretsmanager_secret_rotation.json
var awsSecretsmanagerSecretRotation []byte

//go:embed mapping/aws/resource/states/aws_sfn_alias.json
var awsSfnAlias []byte

//go:embed mapping/aws/resource/aoss/aws_opensearchserverless_access_policy.json
var awsOpenseachserverlessAccessPolicy []byte

//go:embed mapping/aws/resource/aoss/aws_opensearchserverless_collection.json
var awsOpenseachserverlessCollection []byte

//go:embed mapping/aws/resource/aoss/aws_opensearchserverless_lifecycle_policy.json
var awsOpenseachserverlessLifecyclePolicy []byte

//go:embed mapping/aws/resource/aoss/aws_opensearchserverless_security_config.json
var awsOpenseachserverlessSecurityConfig []byte

//go:embed mapping/aws/resource/aoss/aws_opensearchserverless_security_policy.json
var awsOpenseachserverlessSecurityPolicy []byte

//go:embed mapping/aws/resource/aoss/aws_opensearchserverless_vpc_endpoint.json
var awsOpenseachserverlessVpcEndpoint []byte

//go:embed mapping/aws/resource/cloudsearch/aws_cloudsearch_domain.json
var awsCloudsearchDomain []byte

//go:embed mapping/aws/resource/cloudsearch/aws_cloudsearch_domain_service_access_policy.json
var awsCloudsearchDomainServiceAccessPolicy []byte

//go:embed mapping/aws/resource/cloudtrail/aws_cloudtrail_event_data_store.json
var awsCloudtrailEventDataStore []byte

//go:embed mapping/aws/resource/events/aws_cloudwatch_event_bus_policy.json
var awsCloudwatchEventBusPolicy []byte

//go:embed mapping/aws/resource/events/aws_cloudwatch_event_endpoint.json
var awsCloudwatchEventEndpoint []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_data_protection_policy.json
var awsCloudwatchLogDataProtectionPolicy []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_destination.json
var awsCloudwatchLogDestination []byte

//go:embed mapping/aws/resource/kinesis/aws_kinesis_stream_consumer.json
var awsKinesisStreamConsumer []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_destination_policy.json
var awsCloudwatchLogDestinationPolicy []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_destination.json
var awsCloudwatchQueryDestination []byte

//go:embed mapping/aws/resource/datapipeline/aws_datapipeline_pipeline.json
var awsDatapipelinePipeline []byte

//go:embed mapping/aws/resource/datapipeline/aws_datapipeline_pipeline_definition.json
var awsDatapipelinePipelineDefinition []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_proxy_protocol_policy.json
var awsProxyProtocolPolicy []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_app_cookie_stickiness_policy.json
var awsAppCookieStickinessPolicy []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_cache_policy.json
var awsCloudfrontCachePolicy []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_continuous_deployment_policy.json
var awsCloudfrontContinuousDeploymentPolicy []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_origin_request_policy.json
var awsCloudfrontOriginRequestPolicy []byte

//go:embed mapping/aws/resource/codebuild/aws_codebuild_report_group.json
var awsCodebuildReportGroup []byte

//go:embed mapping/aws/resource/codebuild/aws_codebuild_resource_policy.json
var awsCodebuildResourcePolicy []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_registry_policy.json
var awsEcrRegistryPolicy []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_replication_configuration.json
var awsEcrReplicationConfiguration []byte

//go:embed mapping/aws/resource/ecr-public/aws_ecrpublic_repository_policy.json
var awsEcrpublicRepositoryPolicy []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_managed_scaling_policy.json
var awsEmrManagedScalingPolicy []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_cookie_stickiness_policy.json
var awsLbCookieStickinessPolicy []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_byte_match_set.json
var awsWafregionalByteMatchSet []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_geo_match_set.json
var awsWafregionalGeoMatchSet []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_ipset.json
var awsWafregionalIpset []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_rate_based_rule.json
var awsWafregionalRateBasedRule []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_regex_match_set.json
var awsWafregionalRegexMatchSet []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_regex_pattern_set.json
var awsWafregionalRegexPatternSet []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_rule.json
var awsWafregionalRule []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_rule_group.json
var awsWafregionalRuleGroup []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_size_constraint_set.json
var awsWafregionalSizeConstraintSet []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_sql_injection_match_set.json
var awsWafregionalSqlInjectionMatchSet []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_web_acl.json
var awsWafregionalWebAcl []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_xss_match_set.json
var awsWafregionalXssNatchSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_byte_match_set.json
var awsWafByteMatchSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_geo_match_set.json
var awsWafGeoMatchSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_ipset.json
var awsWafIpset []byte

//go:embed mapping/aws/resource/waf/aws_waf_rate_based_rule.json
var awsWafRateBasedRule []byte

//go:embed mapping/aws/resource/waf/aws_waf_regex_match_set.json
var awsWafRegexMatchSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_regex_pattern_set.json
var awsWafRegexPatternSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_rule.json
var awsWafRule []byte

//go:embed mapping/aws/resource/waf/aws_waf_rule_group.json
var awsWafRuleGroup []byte

//go:embed mapping/aws/resource/waf/aws_waf_size_constraint_set.json
var awsWafSizeConstraintSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_sql_injection_match_set.json
var awsWafSqlInjectionMatchSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_web_acl.json
var awsWafWebAcl []byte

//go:embed mapping/aws/resource/waf/aws_waf_xss_match_set.json
var awsWafXssNatchSet []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam.json
var awsVpcIpam []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_pool.json
var awsVpcIpamPool []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_pool_cidr.json
var awsVpcIpamPoolCidr []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_pool_cidr_allocation.json
var awsVpcIpamPoolCidrAllocation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_preview_next_cidr.json
var awsVpcIpamPreviewNextCidr []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_resource_discovery.json
var awsVpcIpamResourceDiscovery []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_resource_discovery_association.json
var awsVpcIpamResourceDiscoveryAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_scope.json
var awsVpcIpamScope []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_security_group_egress_rule.json
var awsVpcSecurityGroupEgressRule []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_security_group_ingress_rule.json
var awsVpcSecurityGroupIngressRule []byte

//go:embed mapping/aws/resource/geo/aws_location_geofence_collection.json
var awsLocationGeofenceCollection []byte

//go:embed mapping/aws/resource/geo/aws_location_map.json
var awsLocationMap []byte

//go:embed mapping/aws/resource/geo/aws_location_place_index.json
var awsLocationPlaceIndex []byte

//go:embed mapping/aws/resource/geo/aws_location_route_calculator.json
var awsLocationRouteCalculator []byte

//go:embed mapping/aws/resource/geo/aws_location_tracker.json
var awsLocationTracker []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_load_balancer_backend_server_policy.json
var awsLoadBalancerBackendServerPolicy []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_load_balancer_listener_policy.json
var awsLoadBalancerListenerPolicy []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_load_balancer_policy.json
var awsLoadBalancerPolicy []byte

//go:embed mapping/aws/resource/route53/aws_route53_cidr_collection.json
var awsRoute53CidrCollection []byte

//go:embed mapping/aws/resource/route53/aws_route53_cidr_location.json
var awsRoute53CidrLocation []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_config.json
var awsRoute53ResolverConfig []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_dnssec_config.json
var awsRoute53ResolverDnssecConfig []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_endpoint.json
var awsRoute53ResolverEndpoint []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_firewall_config.json
var awsRoute53ResolverFirewallConfig []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_firewall_domain_list.json
var awsRoute53ResolverFirewallDomainList []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_firewall_rule.json
var awsRoute53ResolverFirewallRule []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_firewall_rule_group.json
var awsRoute53ResolverFirewallRuleGroup []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_firewall_rule_group_association.json
var awsRoute53ResolverFirewallRuleGroupAssociation []byte

//go:embed mapping/aws/resource/route53resolver/aws_route53_resolver_rule.json
var awsRoute53ResolverRule []byte

//go:embed mapping/aws/resource/route53/aws_route53_traffic_policy.json
var awsRoute53TrafficPolicy []byte

//go:embed mapping/aws/resource/route53/aws_route53_traffic_policy_instance.json
var awsRoute53TrafficPolicyInstance []byte

//go:embed mapping/aws/resource/ce/aws_ce_anomaly_monitor.json
var awsCeAnomalyMonitor []byte

//go:embed mapping/aws/resource/ce/aws_ce_anomaly_subscription.json
var awsCeAnomalySubscription []byte

//go:embed mapping/aws/resource/ce/aws_ce_cost_allocation_tag.json
var awsCeCostAllocationTag []byte

//go:embed mapping/aws/resource/ce/aws_ce_cost_category.json
var awsCeCostCategory []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_function.json
var awsCloudfrontFunction []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_identity_pool.json
var awsCognitoIdentityPool []byte

//go:embed mapping/aws/resource/timestreamwrite/aws_timestreamwrite_database.json
var awsTimestreamwriteDatabase []byte

//go:embed mapping/aws/resource/timestreamwrite/aws_timestreamwrite_table.json
var awsTimestreamwriteTable []byte

//go:embed mapping/aws/resource/codebuild/aws_codebuild_source_credential.json
var awsCodebuildSourceCredential []byte

//go:embed mapping/aws/resource/codecommit/aws_codecommit_approval_rule_template.json
var awsCodecommitApprovalRuleTemplate []byte

//go:embed mapping/aws/resource/codecommit/aws_codecommit_approval_rule_template_association.json
var awsCodecommitApprovalRuleTemplateAssociation []byte

//go:embed mapping/aws/resource/codecommit/aws_codecommit_trigger.json
var awsCodecommitTrigger []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_default_kms_key.json
var awsEbsDefaultKmsKey []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_encryption_by_default.json
var awsEbsEncryptionByDefault []byte

//go:embed mapping/aws/resource/kms/aws_kms_key_policy.json
var awsKmsKeyPolicy []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_availability_zone_group.json
var awsEc2AvailabilityGroup []byte

//go:embed mapping/aws/resource/ec2/aws_eip_association.json
var awsEipAssociation []byte

//go:embed mapping/aws/resource/iam/aws_iam_security_token_service_preferences.json
var awsIamSecurityTokenServicePreferences []byte

//go:embed mapping/aws/resource/iam/aws_iam_service_specific_credential.json
var awsIamServiceSpecificCredential []byte

//go:embed mapping/aws/resource/iam/aws_iam_signing_certificate.json
var awsIamSigningCertificate []byte

//go:embed mapping/aws/resource/iam/aws_iam_virtual_mfa_device.json
var awsIamVirtualMfaDevice []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_container_recipe.json
var awsImagebuilderContainerRecipe []byte

//go:embed mapping/aws/resource/inspector2/aws_inspector2_delegated_admin_account.json
var awsInspector2DelegatedAdminAccount []byte

//go:embed mapping/aws/resource/inspector2/aws_inspector2_enabler.json
var awsInspector2Enabler []byte

//go:embed mapping/aws/resource/inspector2/aws_inspector2_member_association.json
var awsInspector2MemberAssociation []byte

//go:embed mapping/aws/resource/inspector2/aws_inspector2_organization_configuration.json
var awsInspector2OrganizationConfiguration []byte

//go:embed mapping/aws/resource/internetmonitor/aws_internetmonitor_monitor.json
var awsInternetmonitorMonitor []byte

//go:embed mapping/aws/resource/kms/aws_kms_ciphertext.json
var awsKmsCiphertext []byte

//go:embed mapping/aws/resource/kms/aws_kms_custom_key_store.json
var awsKmsCustomKeyStore []byte

//go:embed mapping/aws/resource/kms/aws_kms_external_key.json
var awsKmsExternalKey []byte

//go:embed mapping/aws/resource/kms/aws_kms_replica_external_key.json
var awsKmsReplicaExternalKey []byte

//go:embed mapping/aws/resource/kms/aws_kms_replica_key.json
var awsKmsReplicaKey []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_code_signing_config.json
var awsLambdaCodeSigningConfig []byte

//go:embed mapping/aws/resource/signer/aws_signer_signing_job.json
var awsSignerSigningJob []byte

//go:embed mapping/aws/resource/signer/aws_signer_signing_profile.json
var awsSignerSigningProfile []byte

//go:embed mapping/aws/resource/signer/aws_signer_signing_profile_permission.json
var awsSignerSigningProfilePermission []byte

//go:embed mapping/aws/resource/elasticbeanstalk/aws_elastic_beanstalk_environment.json
var awsElastiBeanstalkEnvironment []byte

//go:embed mapping/aws/resource/elasticbeanstalk/aws_elastic_beanstalk_application_version.json
var awsElasticBeanstalkApplicationVersion []byte

//go:embed mapping/aws/resource/elasticbeanstalk/aws_elastic_beanstalk_configuration_template.json
var awsElasticBeanstalkConfigurationTemplate []byte

//go:embed mapping/aws/resource/codebuild/aws_codebuild_webhook.json
var awsCodebuildWebhook []byte

//go:embed mapping/aws/resource/access-analyzer/aws_accessanalyzer_analyzer.json
var awsAccessAnalyzer []byte

//go:embed mapping/aws/resource/geo/aws_location_tracker_association.json
var awsLocationTrackerAssociation []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_access_log_subscription.json
var awsVpclatticeAccesLogSubscription []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_auth_policy.json
var awsVpclatticeAuthPolicy []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_listener.json
var awsVpclatticeListener []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_listener_rule.json
var awsVpclatticeListernerRule []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_resource_policy.json
var awsVpclatticeResourcePolicy []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_service.json
var awsVpclatticeService []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_service_network.json
var awsVpclatticeServiceNetwork []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_service_network_service_association.json
var awsVpclatticeServiceNetworkServiceAssocation []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_service_network_vpc_association.json
var awsVpclatticeServiceNetworkVpcAssociation []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_target_group.json
var awsVpclatticeTargetGroup []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_target_group_attachment.json
var awsVpclatticeTargetGroupAssociation []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_group_tag.json
var awsAutoscalingGroupTag []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_schedule.json
var awsAutoscalingSchedule []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_traffic_source_attachment.json
var awsAutoscalingTrafficSourceAttachment []byte

//go:embed mapping/aws/resource/autoscaling-plans/aws_autoscalingplans_scaling_plan.json
var awsAutoscalingplansScalingPlan []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_elb_attachment.json
var awsElbAttachment []byte

//go:embed mapping/aws/resource/connect/aws_connect_bot_association.json
var awsConnectBotAssociation []byte

//go:embed mapping/aws/resource/connect/aws_connect_contact_flow.json
var awsConnectContactFlow []byte

//go:embed mapping/aws/resource/connect/aws_connect_contact_flow_module.json
var awsConnectContactFlowModule []byte

//go:embed mapping/aws/resource/connect/aws_connect_hours_of_operation.json
var awsConnectHoursOfOperation []byte

//go:embed mapping/aws/resource/connect/aws_connect_instance.json
var awsConnectInstance []byte

//go:embed mapping/aws/resource/connect/aws_connect_instance_storage_config.json
var awsConnectInstanceStorageConfig []byte

//go:embed mapping/aws/resource/connect/aws_connect_phone_number.json
var awsConnectPhoneNumber []byte

//go:embed mapping/aws/resource/connect/aws_connect_queue.json
var awsConnectQueue []byte

//go:embed mapping/aws/resource/connect/aws_connect_quick_connect.json
var awsConnectQuickConnect []byte

//go:embed mapping/aws/resource/connect/aws_connect_routing_profile.json
var awsConnectRoutingProfile []byte

//go:embed mapping/aws/resource/connect/aws_connect_security_profile.json
var awsConnectSecurityProfile []byte

//go:embed mapping/aws/resource/connect/aws_connect_user.json
var awsConnectUser []byte

//go:embed mapping/aws/resource/connect/aws_connect_user_hierarchy_structure.json
var awsConnectUserHierarchyStructure []byte

//go:embed mapping/aws/resource/connect/aws_connect_user_hierarchy_group.json
var awsConnectUserHierarchyGroup []byte

//go:embed mapping/aws/resource/connect/aws_connect_vocabulary.json
var awsConnectVocabulary []byte

//go:embed mapping/aws/resource/ec2/aws_verifiedaccess_endpoint.json
var awsVerifiedaccessEndpoint []byte

//go:embed mapping/aws/resource/ec2/aws_verifiedaccess_group.json
var awsVerifiedaccessGroup []byte

//go:embed mapping/aws/resource/ec2/aws_verifiedaccess_instance.json
var awsVerifiedaccessInstance []byte

//go:embed mapping/aws/resource/ec2/aws_verifiedaccess_instance_logging_configuration.json
var awsVerifiedaccessInstanceLoggingConfiguration []byte

//go:embed mapping/aws/resource/ec2/aws_verifiedaccess_trust_provider.json
var awsVerifiedaccessTrustProvider []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_budget_resource_association.json
var awsServicecatalogBudgetResourceAssociation []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_constraint.json
var awsServicecatalogConstraint []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_organizations_access.json
var awsServicecatalogOrganizationAccess []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_portfolio_share.json
var awsServicecatalogPortfolioShare []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_principal_portfolio_association.json
var awsServicecatalogPrincipalPortfolioAssociation []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_product.json
var awsServicecatalogProduct []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_product_portfolio_association.json
var awsServicecatalogProductPortfolioAssociation []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_provisioned_product.json
var awsServicecatalogProvisionedProduct []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_service_action.json
var awsServicecatalogServiceAction []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_tag_option.json
var awsServicecatalogTagOption []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_tag_option_resource_association.json
var awsServicecatalogTagOptionResourceAssociation []byte

//go:embed mapping/aws/resource/servicequotas/aws_servicequotas_template.json
var awsServiceQuotasTemplate []byte

//go:embed mapping/aws/resource/servicequotas/aws_servicequotas_template_association.json
var awsServiceQuotasTemplateAssociation []byte

//go:embed mapping/aws/resource/codeguru-profiler/aws_codeguruprofiler_profiling_group.json
var awsCodeguruprofilerProfilingGroup []byte

//go:embed mapping/aws/resource/codeguru-reviewer/aws_codegurureviewer_repository_association.json
var awsCodegurureviewerRepositoryAssociation []byte

//go:embed mapping/aws/resource/codepipeline/aws_codepipeline_custom_action_type.json
var awsCodepipelineCustomActionType []byte

//go:embed mapping/aws/resource/codepipeline/aws_codepipeline_webhook.json
var awsCodepipelineWebhook []byte

//go:embed mapping/aws/resource/codestar-connections/aws_codestarconnections_connection.json
var awsCodestarconnectionsConnection []byte

//go:embed mapping/aws/resource/codestar-connections/aws_codestarconnections_host.json
var awsCodestarconnectionsHost []byte

//go:embed mapping/aws/resource/codestar-notifications/aws_codestarnotifications_notification_rule.json
var awsCodestarconnectionsNotificationsRule []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_account_registration.json
var awsAuditmanagerAccountRegistration []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_assessment.json
var awsAuditmanagerAssessment []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_assessment_delegation.json
var awsAuditmanagerAssessmentDelegation []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_assessment_report.json
var awsAuditmanagerAssessmentReport []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_control.json
var awsAuditmanagerControl []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_framework.json
var awsAuditmanagerFramework []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_framework_share.json
var awsAuditmanagerFrameworkShare []byte

//go:embed mapping/aws/resource/auditmanager/aws_auditmanager_organization_admin_account_registration.json
var awsAuditmanagerOrganizationAdminAccountRegistration []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrock_custom_model.json
var awsBedrockCustomModel []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrock_model_invocation_logging_configuration.json
var awsBedrockModelInvocationLoggingConfiguration []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrock_provisioned_model_throughput.json
var awsBedrockProvisionedModelThroughput []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_agent.json
var awsBedrockagentAgent []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_agent_action_group.json
var awsBedrockagentActionGroup []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_agent_alias.json
var awsBedrockagentAgentAlias []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_agent.json
var awsDatasyncAgent []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_azure_blob.json
var awsDatasyncLocationAzureBlob []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_efs.json
var awsDatasyncLocationEfs []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_fsx_lustre_file_system.json
var awsDatasyncLocationFsxLustreFileSystem []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_fsx_ontap_file_system.json
var awsDatasyncLocationFsxOntapFileSystem []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_fsx_openzfs_file_system.json
var awsDatasyncLocationFsxOpenzfsFileSystem []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_fsx_windows_file_system.json
var awsDatasyncLocationFsxWindowsFileSystem []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_hdfs.json
var awsDatasyncLocationHdfs []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_nfs.json
var awsDatasyncLocationNfs []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_object_storage.json
var awsDatasyncLocationObjectStorage []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_s3.json
var awsDatasyncLocationS3 []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_location_smb.json
var awsDatasyncLocationSmb []byte

//go:embed mapping/aws/resource/datasync/aws_datasync_task.json
var awsDatasyncTask []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_lustre_file_system.json
var awsFsxLustreFileSystem []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_windows_file_system.json
var awsFsxWindowsFileSystem []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_backup.json
var awsFsxBackup []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_data_repository_association.json
var awsFsxDataRepositoryAssociation []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_file_cache.json
var awsFsxFileCache []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_ontap_file_system.json
var awsFsxOntapFileSystem []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_ontap_storage_virtual_machine.json
var awsFsxOntapStorageVirtualMachine []byte

//go:embed mapping/aws/resource/fsx/aws_fsx_ontap_volume.json
var awsFsxOntapVolume []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_domain.json
var awsDatazoneDomain []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_environment_blueprint_configuration.json
var awsDatazoneEnvironmentBlueprintConfiguration []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_cache.json
var awsStoragegatewayCache []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_cached_iscsi_volume.json
var awsStoragegatewayCachedIscsiVolume []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_file_system_association.json
var awsStoragegatewayFileSystemAssociation []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_gateway.json
var awsStoragegatewayGateway []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_nfs_file_share.json
var awsStoragegatewayNfsFileShare []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_smb_file_share.json
var awsStoragegatewaySmbFileShare []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_stored_iscsi_volume.json
var awsStoragegatewayStoredIscsiVolume []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_tape_pool.json
var awsStoragegatewayTapePool []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_upload_buffer.json
var awsStoragegatewayUploadBuffer []byte

//go:embed mapping/aws/resource/storagegateway/aws_storagegateway_working_storage.json
var awsStoragegatewayWorkingStorage []byte

//go:embed mapping/aws/resource/ses/aws_ses_active_receipt_rule_set.json
var awsSesActiveReceiptRuleSet []byte

//go:embed mapping/aws/resource/ses/aws_ses_configuration_set.json
var awsSesConfigurationSet []byte

//go:embed mapping/aws/resource/ses/aws_ses_email_identity.json
var awsSesEmailIdentity []byte

//go:embed mapping/aws/resource/ses/aws_ses_event_destination.json
var awsSesEventDestination []byte

//go:embed mapping/aws/resource/ses/aws_ses_receipt_filter.json
var awsSesReceiptFilter []byte

//go:embed mapping/aws/resource/ses/aws_ses_template.json
var awsSesTemplate []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_app.json
var awsSagemakerApp []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_app_image_config.json
var awsSagemakerAppImageConfig []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_code_repository.json
var awsSagemakerCodeRepository []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_data_quality_job_definition.json
var awsSagemakerDataQualityJobDefinition []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_device.json
var awsSagemakerDevice []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_device_fleet.json
var awsSagemakerDeviceFleet []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_domain.json
var awsSagemakerDomain []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_endpoint.json
var awsSagemakerEndpoint []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_feature_group.json
var awsSagemakerFeatureGroup []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_image.json
var awsSagemakerImage []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_image_version.json
var awsSagemakerImageVersion []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_model_package_group.json
var awsSagemakerModelPackageGroup []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_model_package_group_policy.json
var awsSagemakerModelPackageGroupPolicy []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_monitoring_schedule.json
var awsSagemakerMonitoringSchedule []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_notebook_instance.json
var awsSagemakerNotebookInstance []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_notebook_instance_lifecycle_configuration.json
var awsSagemakerNotebookInstanceLifecycleConfiguration []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_pipeline.json
var awsSagemakerPipeline []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_project.json
var awsSagemakerProject []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_servicecatalog_portfolio_status.json
var awsSagemakerServicecatalogPortfolioStatus []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_space.json
var awsSagemakerSpace []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_studio_lifecycle_config.json
var awsSagemakerStudioLifecycleConfig []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_user_profile.json
var awsSagemakerUserProfile []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_workforce.json
var awsSagemakerWorkforce []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_workteam.json
var awsSagemakerWorkteam []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_carrier_gateway.json
var awsEc2CarrierGateway []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_client_vpn_authorization_rule.json
var awsEc2ClientvpnAuthorizationRule []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_client_vpn_endpoint.json
var awsEc2ClientVpnEndpoint []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_client_vpn_network_association.json
var awsEc2ClientVpnNetworkAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_client_vpn_route.json
var awsEc2ClientVpnRoute []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_fleet.json
var awsEc2Fleet []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_host.json
var awsEc2Host []byte

//go:embed mapping/aws/resource/access-analyzer/aws_accessanalyzer_archive_rule.json
var awsAccessAnalyzerArchiveRule []byte

//go:embed mapping/aws/resource/account/aws_account_region.json
var awsAccountRegion []byte

//go:embed mapping/aws/resource/acm-pa/aws_acmpca_permission.json
var awsAcmpcaPermission []byte

//go:embed mapping/aws/resource/acm-pa/aws_acmpca_policy.json
var awsAcmpcaPolicy []byte

//go:embed mapping/aws/resource/sdb/aws_simpledb_domain.json
var awsSimpledbDomain []byte

//go:embed mapping/aws/resource/ec2/aws_snapshot_create_volume_permission.json
var awsSnapshotCreateVolumePermission []byte

//go:embed mapping/aws/resource/sns/aws_sns_platform_application.json
var awsSnsPlatformApplication []byte

//go:embed mapping/aws/resource/synthetics/aws_synthetics_canary.json
var awsSyntheticsCanary []byte

//go:embed mapping/aws/resource/synthetics/aws_synthetics_group.json
var awsSyntheticsGroup []byte

//go:embed mapping/aws/resource/synthetics/aws_synthetics_group_association.json
var awsSyntheticsGroupAssociation []byte

//go:embed mapping/aws/resource/dms/aws_dms_certificate.json
var awsDmsCertificate []byte

//go:embed mapping/aws/resource/dms/aws_dms_endpoint.json
var awsDmsEndpoint []byte

//go:embed mapping/aws/resource/dms/aws_dms_event_subscription.json
var awsDmsEventSubscription []byte

//go:embed mapping/aws/resource/dms/aws_dms_replication_instance.json
var awsDmsReplicationInstance []byte

//go:embed mapping/aws/resource/dms/aws_dms_replication_subnet_group.json
var awsDmsReplicationSubnetGroup []byte

//go:embed mapping/aws/resource/dms/aws_dms_replication_task.json
var awsDmsReplicationTask []byte

//go:embed mapping/aws/resource/dms/aws_dms_s3_endpoint.json
var awsDmsS3Endpoint []byte

//go:embed mapping/aws/resource/dms/aws_dms_replication_config.json
var awsDmsReplicationConfig []byte

//go:embed mapping/aws/resource/rds/aws_db_instance_automated_backups_replication.json
var awsDbInstanceAutomatedBackupsReplication []byte

//go:embed mapping/aws/resource/rds/aws_db_instance_role_association.json
var awsDbInstanceRoleAssociation []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy.json
var awsDbProxy []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy_default_target_group.json
var awsDbProxyDefaultTargetGroup []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy_endpoint.json
var awsDbProxyEndpoint []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy_target.json
var awsDbProxyTarget []byte

//go:embed mapping/aws/resource/rds/aws_db_snapshot.json
var awsDbSnapshot []byte

//go:embed mapping/aws/resource/rds/aws_db_snapshot_copy.json
var awsDbSnapshotCopy []byte

//go:embed mapping/aws/resource/transcribe/aws_transcribe_language_model.json
var awsTranscribeLanguageModel []byte

//go:embed mapping/aws/resource/transcribe/aws_transcribe_medical_vocabulary.json
var awsTranscribeMedicalVocabulary []byte

//go:embed mapping/aws/resource/transcribe/aws_transcribe_vocabulary.json
var awsTranscribeVocabulary []byte

//go:embed mapping/aws/resource/transcribe/aws_transcribe_vocabulary_filter.json
var awsTranscribeVocabularyFilter []byte

//go:embed mapping/aws/resource/oam/aws_oam_link.json
var awsOamLink []byte

//go:embed mapping/aws/resource/oam/aws_oam_sink.json
var awsOamSink []byte

//go:embed mapping/aws/resource/oam/aws_oam_sink_policy.json
var awsOamSinkPolicy []byte
