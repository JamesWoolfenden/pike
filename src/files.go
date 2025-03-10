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
var awsWafregionalSQLInjectionMatchSet []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_web_acl.json
var awsWafregionalWebACL []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_xss_match_set.json
var awsWafregionalXSSNatchSet []byte

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
var awsWafSQLInjectionMatchSet []byte

//go:embed mapping/aws/resource/waf/aws_waf_web_acl.json
var awsWafWebACL []byte

//go:embed mapping/aws/resource/waf/aws_waf_xss_match_set.json
var awsWafXSSNatchSet []byte

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
var awsDBInstanceAutomatedBackupsReplication []byte

//go:embed mapping/aws/resource/rds/aws_db_instance_role_association.json
var awsDBInstanceRoleAssociation []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy.json
var awsDBProxy []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy_default_target_group.json
var awsDBProxyDefaultTargetGroup []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy_endpoint.json
var awsDBProxyEndpoint []byte

//go:embed mapping/aws/resource/rds/aws_db_proxy_target.json
var awsDBProxyTarget []byte

//go:embed mapping/aws/resource/rds/aws_db_snapshot.json
var awsDBSnapshot []byte

//go:embed mapping/aws/resource/rds/aws_db_snapshot_copy.json
var awsDBSnapshotCopy []byte

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

//go:embed mapping/aws/resource/amplify/aws_amplify_app.json
var awsAmplifyApp []byte

//go:embed mapping/aws/resource/amplify/aws_amplify_branch.json
var awsAmplifyBranch []byte

//go:embed mapping/aws/resource/amplify/aws_amplify_domain_association.json
var awsAmplifyDomainAssociation []byte

//go:embed mapping/aws/resource/workspaces/aws_workspaces_connection_alias.json
var awsWorkspacesConnectionAlias []byte

//go:embed mapping/aws/resource/workspaces/aws_workspaces_ip_group.json
var awsWorkspacesIPGroup []byte

//go:embed mapping/aws/resource/chime/aws_chime_voice_connector.json
var awsChimeVoiceConnector []byte

//go:embed mapping/aws/resource/chime/aws_chime_voice_connector_group.json
var awsChimeVoiceConnectorGroup []byte

//go:embed mapping/aws/resource/chime/aws_chime_voice_connector_logging.json
var awsChimeVoiceConnectorLogging []byte

//go:embed mapping/aws/resource/chime/aws_chime_voice_connector_origination.json
var awsChimeVoiceConnectorOrigination []byte

//go:embed mapping/aws/resource/chime/aws_chime_voice_connector_streaming.json
var awsChimeVoiceConnectorStreaming []byte

//go:embed mapping/aws/resource/chime/aws_chime_voice_connector_termination.json
var awsChimeVoiceConnectorTermination []byte

//go:embed mapping/aws/resource/chime/aws_chime_voice_connector_termination_credentials.json
var awsChimeVoiceConnectorTerminationCredentials []byte

//go:embed mapping/aws/resource/chime/aws_chimesdkmediapipelines_media_insights_pipeline_configuration.json
var awsChimesdkmediapipelinesMediaInsightsPipelineConfiguration []byte

//go:embed mapping/aws/resource/chime/aws_chimesdkvoice_global_settings.json
var awsChimesdkvoiceGlobalSettings []byte

//go:embed mapping/aws/resource/chime/aws_chimesdkvoice_sip_media_application.json
var awsChimesdkvoiceSIPMediaApplication []byte

//go:embed mapping/aws/resource/chime/aws_chimesdkvoice_sip_rule.json
var awsChimesdkvoiceSIPRule []byte

//go:embed mapping/aws/resource/chime/aws_chimesdkvoice_voice_profile_domain.json
var awsChimesdkvoiceVoiceProfileDomain []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_environment.json
var awsAppconfigEnvironment []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_extension.json
var awsAppconfigExtension []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_extension_association.json
var awsAppconfigExtensionAssociation []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_data_source.json
var awsBedrockagentDatasource []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_knowledge_base.json
var awsBedrockagentKnowledgeBase []byte

//go:embed mapping/aws/resource/cleanrooms/aws_cleanrooms_collaboration.json
var awsCleanroomsCollaboration []byte

//go:embed mapping/aws/resource/cleanrooms/aws_cleanrooms_configured_table.json
var awsCleanroomsConfiguredTable []byte

//go:embed mapping/aws/resource/ec2/aws_vpn_connection_route.json
var awsVpnConnectionRoute []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_hosted_configuration_version.json
var awsAppconfigHostedConfigurationVersion []byte

//go:embed mapping/aws/resource/appflow/aws_appflow_connector_profile.json
var awsAppflowConnectorProfile []byte

//go:embed mapping/aws/resource/appflow/aws_appflow_flow.json
var awsAppflowFlow []byte

//go:embed mapping/aws/resource/app-integrations/aws_appintegrations_data_integration.json
var awsAppintegrationsDataIntegration []byte

//go:embed mapping/aws/resource/app-integrations/aws_appintegrations_event_integration.json
var awsAppintegrationsEventIntegration []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_auto_scaling_configuration_version.json
var awsApprunnerAutoScalingConfigurationVersion []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_default_auto_scaling_configuration_version.json
var awsApprunnerDefaultAutoScalingConfigurationVersion []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_observability_configuration.json
var awsApprunnerObservabilityConfiguration []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_service.json
var awsApprunnerService []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_vpc_connector.json
var awsApprunnerVpcConnector []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_vpc_ingress_connection.json
var awsApprunnerVpcIngressConnection []byte

//go:embed mapping/aws/resource/appstream/aws_appstream_image_builder.json
var awsAppstreamImageBuilder []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_domain_name.json
var awsAppsyncDomainName []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_domain_name_api_association.json
var awsAppsyncDomainNameAPIAssociation []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_function.json
var awsAppsyncFunction []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_resolver.json
var awsAppsyncResolver []byte

//go:embed mapping/aws/resource/athena/aws_athena_prepared_statement.json
var awsAthenaPreparedStatement []byte

//go:embed mapping/aws/resource/bcm-data-exports/aws_bcmdataexports_export.json
var awsBcmdataexportsExport []byte

//go:embed mapping/aws/resource/chatbot/aws_chatbot_slack_channel_configuration.json
var awsChatbotSlackChannelConfiguration []byte

//go:embed mapping/aws/resource/chatbot/aws_chatbot_teams_channel_configuration.json
var awsChatbotTeamsChannelConfiguration []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_key_value_store.json
var awsCloudfrontKeyValueStore []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_realtime_log_config.json
var awsCloudfrontRealtimeLogConfig []byte

//go:embed mapping/aws/resource/cassandra/aws_keyspaces_keyspace.json
var awsKeyspacesKeyspace []byte

//go:embed mapping/aws/resource/cassandra/aws_keyspaces_table.json
var awsKeyspacesTable []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_identity_pool_provider_principal_tag.json
var awsCognitoIdentityPoolProviderPrincipalTag []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_identity_pool_roles_attachment.json
var awsCognitoIdentityPoolRolesAttachment []byte

//go:embed mapping/aws/resource/comprehend/aws_comprehend_document_classifier.json
var awsComprehendDocumentClassifier []byte

//go:embed mapping/aws/resource/config/aws_config_conformance_pack.json
var awsConfigConformancePack []byte

//go:embed mapping/aws/resource/config/aws_config_organization_conformance_pack.json
var awsConfigOrganizationConformancePack []byte

//go:embed mapping/aws/resource/controltower/aws_controltower_control.json
var awsControltowerControl []byte

//go:embed mapping/aws/resource/controltower/aws_controltower_landing_zone.json
var awsControltowerLandingZone []byte

//go:embed mapping/aws/resource/cur/aws_cur_report_definition.json
var awsCurReportDefinition []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_project.json
var awsDatazoneProject []byte

//go:embed mapping/aws/resource/detective/aws_detective_graph.json
var awsDetectiveGraph []byte

//go:embed mapping/aws/resource/detective/aws_detective_invitation_accepter.json
var awsDetectiveInvitationAccepter []byte

//go:embed mapping/aws/resource/detective/aws_detective_organization_admin_account.json
var awsDectectiveOrganizationAdminAccount []byte

//go:embed mapping/aws/resource/devops-guru/aws_devopsguru_notification_channel.json
var awsDevopsguruNotificationChannel []byte

//go:embed mapping/aws/resource/devops-guru/aws_devopsguru_resource_collection.json
var awsDevopsguruResourceCollection []byte

//go:embed mapping/aws/resource/docdb-elastic/aws_docdbelastic_cluster.json
var awsDocdbelasticCluster []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_local_gateway_route.json
var awsEc2LocalGatewayRoute []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_local_gateway_route_table_vpc_association.json
var awsEc2LocalGatewayRouteTableVpcAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_managed_prefix_list.json
var awsEc2ManagedPrefixList []byte

//go:embed mapping/aws/resource/ec2/aws_network_interface_attachment.json
var awsNetworkInterfaceAttachment []byte

//go:embed mapping/aws/resource/ec2/aws_spot_fleet_request.json
var awsSpotFleetRequest []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_connect.json
var awsEc2TransitGatewayConnect []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_multicast_domain.json
var awsEc2TransitGatewayMulticastDomain []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_multicast_domain_association.json
var awsEc2TransitGatewayMulticastDomainAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_multicast_group_member.json
var awsEc2TransitGatewayMulticastGroupMember []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_multicast_group_source.json
var awsEc2TransitGatewayMulticastGroupSource []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_peering_attachment.json
var awsEc2TransitGatewayPeeringAttachment []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_repository_creation_template.json
var awsEcrRepositoryCreationTemplate []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_capacity_provider.json
var awsEcsCapacityProvider []byte

//go:embed mapping/aws/resource/ec2/aws_network_acl_association.json
var awsNetworkACLAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_connection_notification.json
var awsVpcEndpointConnectionNotification []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_task_set.json
var awsEcsTaskSet []byte

//go:embed mapping/aws/resource/eks/aws_eks_access_entry.json
var awsEksAccessEntry []byte

//go:embed mapping/aws/resource/eks/aws_eks_fargate_profile.json
var awsEksFargateProfile []byte

//go:embed mapping/aws/resource/eks/aws_eks_identity_provider_config.json
var awsEksIdentityProviderConfig []byte

//go:embed mapping/aws/resource/eks/aws_eks_pod_identity_association.json
var awsEksPodIdentityAssociation []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_global_replication_group.json
var awsElasticacheGlobalReplicationGroup []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_studio.json
var awsEmrStudio []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_studio_session_mapping.json
var awsEmrStudioSessionMapping []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emrcontainers_virtual_cluster.json
var awsEmrcontainersVirtualCluster []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emrserverless_application.json
var awsEmrseverlessApplication []byte

//go:embed mapping/aws/resource/evidently/aws_evidently_feature.json
var awsEvidentlyFeature []byte

//go:embed mapping/aws/resource/evidently/aws_evidently_launch.json
var awsEvidentlyLaunch []byte

//go:embed mapping/aws/resource/evidently/aws_evidently_project.json
var awsEvidentlyProject []byte

//go:embed mapping/aws/resource/evidently/aws_evidently_segment.json
var awsEvidentlySegment []byte

//go:embed mapping/aws/resource/schemas/aws_schemas_discoverer.json
var awsSchemasDiscoverer []byte

//go:embed mapping/aws/resource/schemas/aws_schemas_registry.json
var awsSchemasRegistry []byte

//go:embed mapping/aws/resource/schemas/aws_schemas_registry_policy.json
var awsSchemasRegistryPolicy []byte

//go:embed mapping/aws/resource/schemas/aws_schemas_schema.json
var awsSchemasSchema []byte

//go:embed mapping/aws/resource/finspace/aws_finspace_kx_environment.json
var awsFinspaceKxEnvironment []byte

//go:embed mapping/aws/resource/fis/aws_fis_experiment_template.json
var awsFisExperimentTemplate []byte

//go:embed mapping/aws/resource/fms/aws_fms_policy.json
var awsFmsPolicy []byte

//go:embed mapping/aws/resource/fms/aws_fms_resource_set.json
var awsFmsResourceSet []byte

//go:embed mapping/aws/resource/gamelift/aws_gamelift_alias.json
var awsGameliftAlias []byte

//go:embed mapping/aws/resource/gamelift/aws_gamelift_build.json
var awsGameliftBuild []byte

//go:embed mapping/aws/resource/gamelift/aws_gamelift_fleet.json
var awsGameliftFleet []byte

//go:embed mapping/aws/resource/gamelift/aws_gamelift_game_server_group.json
var awsGameliftGameServerGroup []byte

//go:embed mapping/aws/resource/gamelift/aws_gamelift_game_session_queue.json
var awsGameliftGameSessionQueue []byte

//go:embed mapping/aws/resource/gamelift/aws_gamelift_script.json
var awsGameliftScript []byte

//go:embed mapping/aws/resource/globalaccelerator/aws_globalaccelerator_accelerator.json
var awsGlobalacceleratorAccelerator []byte

//go:embed mapping/aws/resource/globalaccelerator/aws_globalaccelerator_cross_account_attachment.json
var awsGlobalacceleratorCrossAccountAttachment []byte

//go:embed mapping/aws/resource/globalaccelerator/aws_globalaccelerator_endpoint_group.json
var awsGlobalacceleratorEndpointGroup []byte

//go:embed mapping/aws/resource/globalaccelerator/aws_globalaccelerator_listener.json
var awsGlobalacceleratorListener []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_detector.json
var awsGuarddutyDetector []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_filter.json
var awsGuarddutyFilter []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_ipset.json
var awsGuarddutyIpset []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_malware_protection_plan.json
var awsGuarddutyMalwareProtectionPlan []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_member.json
var awsGuarddutyMember []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_threatintelset.json
var awsGuarddutyThreatintelset []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_workflow.json
var awsImagebuilderWorkflow []byte

//go:embed mapping/aws/resource/iot/aws_iot_authorizer.json
var awsIotAuthorizer []byte

//go:embed mapping/aws/resource/iot/aws_iot_billing_group.json
var awsIotBillingGroup []byte

//go:embed mapping/aws/resource/iot/aws_iot_ca_certificate.json
var awsIotCaCertificate []byte

//go:embed mapping/aws/resource/iot/aws_iot_certificate.json
var awsIotCertificate []byte

//go:embed mapping/aws/resource/iot/aws_iot_policy.json
var awsIotPolicy []byte

//go:embed mapping/aws/resource/iot/aws_iot_provisioning_template.json
var awsIotProvisioningTemplate []byte

//go:embed mapping/aws/resource/iot/aws_iot_role_alias.json
var awsIotRoleAlias []byte

//go:embed mapping/aws/resource/iot/aws_iot_thing.json
var awsIotThing []byte

//go:embed mapping/aws/resource/iot/aws_iot_thing_group.json
var awsIotThingGroup []byte

//go:embed mapping/aws/resource/iot/aws_iot_thing_type.json
var awsIotThingType []byte

//go:embed mapping/aws/resource/iot/aws_iot_topic_rule.json
var awsIotTopicRule []byte

//go:embed mapping/aws/resource/iot/aws_iot_topic_rule_destination.json
var awsIotTopicRuleDestination []byte

//go:embed mapping/aws/resource/ivschat/aws_ivs_channel.json
var awsIvsChannel []byte

//go:embed mapping/aws/resource/ivschat/aws_ivs_playback_key_pair.json
var awsIvsPlaybackKeyPair []byte

//go:embed mapping/aws/resource/ivschat/aws_ivs_recording_configuration.json
var awsIvsRecordingConfigration []byte

//go:embed mapping/aws/resource/ivschat/aws_ivschat_logging_configuration.json
var awsIvschatLoggingConfiguration []byte

//go:embed mapping/aws/resource/ivschat/aws_ivschat_room.json
var awsIvschatRoom []byte

//go:embed mapping/aws/resource/kendra/aws_kendra_data_source.json
var awsKendraDataSource []byte

//go:embed mapping/aws/resource/kendra/aws_kendra_faq.json
var awsKendraFaq []byte

//go:embed mapping/aws/resource/kendra/aws_kendra_index.json
var awsKendraIndex []byte

//go:embed mapping/aws/resource/kinesisanalytics/aws_kinesisanalyticsv2_application.json
var awsKinesisanalyticsv2Application []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_data_cells_filter.json
var awsLakeformationDataCellsFilter []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_resource_lf_tag.json
var awsLakeformationResourceLfTag []byte

//go:embed mapping/aws/resource/lex/aws_lex_bot.json
var awsLexBot []byte

//go:embed mapping/aws/resource/lex/aws_lex_bot_alias.json
var awsLexBotAlias []byte

//go:embed mapping/aws/resource/lex/aws_lex_intent.json
var awsLexIntent []byte

//go:embed mapping/aws/resource/lex/aws_lex_slot_type.json
var awsLexSlotType []byte

//go:embed mapping/aws/resource/license-manager/aws_licensemanager_grant.json
var awsLicencemanagerGrant []byte

//go:embed mapping/aws/resource/license-manager/aws_licensemanager_license_configuration.json
var awsLicenceManagerLicencenceConfiguration []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_bucket.json
var awsLightsailBucket []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_certificate.json
var awsLightsailCertificate []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_database.json
var awsLightsailDatabase []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_disk.json
var awsLightsailDisk []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_distribution.json
var awsLightsailDistribution []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_lb.json
var awsLightsailLb []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_account_policy.json
var awsCloudwatchLogAccountPolicy []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_custom_data_identifier.json
var awsMacie2CustomDataIdentifier []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_findings_filter.json
var awsMacie2FindingsFilter []byte

//go:embed mapping/aws/resource/mediapackage/aws_media_package_channel.json
var awsMediaPackageChannel []byte

//go:embed mapping/aws/resource/medialive/aws_medialive_multiplex.json
var awsMedialiveMulitplex []byte

//go:embed mapping/aws/resource/medialive/aws_medialive_multiplex_program.json
var awsMedialiveMultiplexProgram []byte

//go:embed mapping/aws/resource/memorydb/aws_memorydb_acl.json
var awsMemorydbACL []byte

//go:embed mapping/aws/resource/memorydb/aws_memorydb_parameter_group.json
var awsMemorydbParameterGroup []byte

//go:embed mapping/aws/resource/kafka/aws_msk_cluster_policy.json
var awsMskClusterPolicy []byte

//go:embed mapping/aws/resource/kafka/aws_msk_replicator.json
var awsMskReplicator []byte

//go:embed mapping/aws/resource/kafka/aws_msk_vpc_connection.json
var awsMskVpcConnection []byte

//go:embed mapping/aws/resource/airflow/aws_mwaa_environment.json
var awsMwaaEnvironment []byte

//go:embed mapping/aws/resource/network-firewall/aws_networkfirewall_tls_inspection_configuration.json
var awsNetworkfirewallTLSInspectionConfiguration []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_connect_attachment.json
var awsNetworkManagerConnectAttachment []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_connect_peer.json
var awsNetworkManagerConnectPeer []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_core_network.json
var awsNetworkmanagerCoreNetwork []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_customer_gateway_association.json
var awsNetworkmanagerCustomerGatewayAssociation []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_device.json
var awsNetworkmanagerDevice []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_global_network.json
var awsNetworkmanagerGlobalNetwork []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_link.json
var awsNetworkmanagerLink []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_link_association.json
var awsNetworkmanagerLinkAssocation []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_site.json
var awsNetworkmanagerSite []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_site_to_site_vpn_attachment.json
var awsNetworkmanagerSiteToSiteVpnAttachment []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_transit_gateway_peering.json
var awsNetworkmanagerTransitGatewayPeering []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_transit_gateway_registration.json
var awsNetworkmanagerTransitGatewayRegistration []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_transit_gateway_route_table_attachment.json
var awsNetworkmanagerTransitGatewayRouteTableAttachment []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_vpc_attachment.json
var awsNetworkmanagerVpcAttachment []byte

//go:embed mapping/aws/resource/organizations/aws_organizations_account.json
var awsOrganizationsAccount []byte

//go:embed mapping/aws/resource/organizations/aws_organizations_organization.json
var awsOrganizationsOrganization []byte

//go:embed mapping/aws/resource/organizations/aws_organizations_organizational_unit.json
var awsOrganizationsOrganizationalUnit []byte

//go:embed mapping/aws/resource/organizations/aws_organizations_resource_policy.json
var awsOrganizationsResourcePolicy []byte

//go:embed mapping/aws/resource/osis/aws_osis_pipeline.json
var awsOsisPipeline []byte

//go:embed mapping/aws/resource/payment-cryptography/aws_paymentcryptography_key.json
var awsPaymentcrytopgraphyKey []byte

//go:embed mapping/aws/resource/payment-cryptography/aws_paymentcryptography_key_alias.json
var awsPaymentcrytopgraphyKeyAlias []byte

//go:embed mapping/aws/resource/pipes/aws_pipes_pipe.json
var awsPipesPipe []byte

//go:embed mapping/aws/resource/qldb/aws_qldb_stream.json
var awsQldbStream []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_analysis.json
var awsQuicksightAnalysis []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_dashboard.json
var awsQuicksightDashboard []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_data_set.json
var awsQuicksightDataSet []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_data_source.json
var awsQuicksightDataSource []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_refresh_schedule.json
var awsQuicksightRefreshSchedule []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_template.json
var awsQuicksightTemplate []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_theme.json
var awsQuicksightTheme []byte

//go:embed mapping/aws/resource/rds/aws_rds_integration.json
var awsRdsIntegration []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_endpoint_access.json
var awsRedshiftEndpointAccess []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_endpoint_authorization.json
var awsRedshiftEndpointAuthorization []byte

//go:embed mapping/aws/resource/redshift-serverless/aws_redshiftserverless_namespace.json
var awsRedshiftserverlessNamespace []byte

//go:embed mapping/aws/resource/redshift-serverless/aws_redshiftserverless_workgroup.json
var awsRedshiftserverlessWorkgroup []byte

//go:embed mapping/aws/resource/rekognition/aws_rekognition_collection.json
var awsRekognitionCollection []byte

//go:embed mapping/aws/resource/rekognition/aws_rekognition_project.json
var awsRekognitionProject []byte

//go:embed mapping/aws/resource/rekognition/aws_rekognition_stream_processor.json
var awsRekognitionStreamProcessor []byte

//go:embed mapping/aws/resource/resource-explorer-2/aws_resourceexplorer2_index.json
var awsResourceexplorer2Index []byte

//go:embed mapping/aws/resource/resource-explorer-2/aws_resourceexplorer2_view.json
var awsResourceexplorer2View []byte

//go:embed mapping/aws/resource/rolesanywhere/aws_rolesanywhere_profile.json
var awsRolesanywhereProfile []byte

//go:embed mapping/aws/resource/rolesanywhere/aws_rolesanywhere_trust_anchor.json
var awsRolesanywhereTrustAnchor []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoverycontrolconfig_cluster.json
var awsRoute53recoverycontrolconfigCluster []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoverycontrolconfig_control_panel.json
var awsRoute53recoverycontrolconfigControlPanel []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoverycontrolconfig_routing_control.json
var awsRoute53recoverycontrolconfigRoutingControl []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoverycontrolconfig_safety_rule.json
var awsRoute53recoverycontrolconfigSafetyRule []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoveryreadiness_cell.json
var awsRoute53recoveryreadinessCell []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoveryreadiness_readiness_check.json
var awsRoute53recoveryreadinessReadinessCheck []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoveryreadiness_recovery_group.json
var awsRoute53recoveryreadinessRecoveryGroup []byte

//go:embed mapping/aws/resource/route53-recovery-readiness/aws_route53recoveryreadiness_resource_set.json
var awsRoute53recoveryreadinessResourceSet []byte

//go:embed mapping/aws/resource/rum/aws_rum_app_monitor.json
var awsRumAppMonitor []byte

//go:embed mapping/aws/resource/s3/aws_s3_access_point.json
var awsS3AccessPoint []byte

//go:embed mapping/aws/resource/s3/aws_s3control_access_grant.json
var awsS3controlAccessGrant []byte

//go:embed mapping/aws/resource/s3/aws_s3control_access_grants_instance.json
var awsS3controlAccessGrantInstance []byte

//go:embed mapping/aws/resource/s3/aws_s3control_access_grants_location.json
var awsS3controlAccessGrantLocation []byte

//go:embed mapping/aws/resource/s3/aws_s3control_multi_region_access_point.json
var awscontrolMultiRegionAccessPoint []byte

//go:embed mapping/aws/resource/s3/aws_s3control_multi_region_access_point_policy.json
var awscontrolMultiRegionAccessPointPolicy []byte

//go:embed mapping/aws/resource/s3-outposts/aws_s3outposts_endpoint.json
var awsS3outpostsEndpoint []byte

//go:embed mapping/aws/resource/scheduler/aws_scheduler_schedule.json
var awsSchedulerSchedule []byte

//go:embed mapping/aws/resource/scheduler/aws_scheduler_schedule_group.json
var awsSchedulerScheduleGroup []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_automation_rule.json
var awsSecurityhubAutomationRule []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_configuration_policy.json
var awsSecurityhubConfigurationPolicy []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_finding_aggregator.json
var awsSecurityhubFindingAggregator []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_permission_set.json
var awsSsoadminPermissionSet []byte

//go:embed mapping/aws/resource/timestreamwrite/aws_timestreaminfluxdb_db_instance.json
var awsTimestreamhubFindingAggregator []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_insight.json
var awsSecurityhubInsight []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_organization_configuration.json
var awsSecurityhubOrganizationConfiguration []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_product_subscription.json
var awsSecurityhubProductSubscription []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_standards_control.json
var awsSecurityhubStandardsControl []byte

//go:embed mapping/aws/resource/securitylake/aws_securitylake_aws_log_source.json
var awsSecuritylakeAwsLogSource []byte

//go:embed mapping/aws/resource/securitylake/aws_securitylake_data_lake.json
var awsSecuritylakeDataLake []byte

//go:embed mapping/aws/resource/securitylake/aws_securitylake_subscriber.json
var awsSecuritylakeSubscriber []byte

//go:embed mapping/aws/resource/securitylake/aws_securitylake_subscriber_notification.json
var awsSecuritylakeSubscriberNotification []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalogappregistry_application.json
var awsServicecatalogappregistryApplication []byte

//go:embed mapping/aws/resource/shield/aws_shield_proactive_engagement.json
var awsShieldProactiveEngagement []byte

//go:embed mapping/aws/resource/shield/aws_shield_protection.json
var awsShieldProtection []byte

//go:embed mapping/aws/resource/shield/aws_shield_protection_group.json
var awsShieldProtectionGroup []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_association.json
var awsSsmAssociation []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_resource_data_sync.json
var awsSsmResourceDataSync []byte

//go:embed mapping/aws/resource/ssm-contacts/aws_ssmcontacts_rotation.json
var awsSsmcontactsRotation []byte

//go:embed mapping/aws/resource/ssm-incidents/aws_ssmincidents_response_plan.json
var awsSsmincidentsResponsePlan []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_application.json
var awsSsoadminApplication []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_application_assignment.json
var awsSsoadminApplicationAssignment []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_agreement.json
var awsTransferAgreement []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_certificate.json
var awsTransferCertificate []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_connector.json
var awsTransferConnector []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_profile.json
var awsTransferProfile []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_workflow.json
var awsTransferWorkflow []byte

//go:embed mapping/aws/resource/verifiedpermissions/aws_verifiedpermissions_identity_source.json
var awsVerifiedpermissionsIdentitySource []byte

//go:embed mapping/aws/resource/verifiedpermissions/aws_verifiedpermissions_policy.json
var awsVerifiedpermissionsPolicy []byte

//go:embed mapping/aws/resource/verifiedpermissions/aws_verifiedpermissions_policy_store.json
var awsVerifiedpermissionsPolicyStore []byte

//go:embed mapping/aws/resource/verifiedpermissions/aws_verifiedpermissions_policy_template.json
var awsVerifiedpermissionsPolicyTemplate []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_snapshot_block_public_access.json
var awsEbsSnapshotBlockPublicAccess []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_image_block_public_access.json
var awsEc2ImageBlockPublicAccess []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_serial_console_access.json
var awsEc2SerialConsoleAccess []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_block_public_access_configuration.json
var awsEmrBlockPublicAccessConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_account_public_access_block.json
var awsS3AccountPublicAccessBlock []byte

//go:embed mapping/aws/resource/s3/aws_s3control_access_point_policy.json
var awsS3controlAccessPointPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_role_policies_exclusive.json
var awsIamGroupPoliciesExclusive []byte

//go:embed mapping/aws/resource/iam/aws_iam_role_policies_exclusive.json
var awsIamRolePoliciesExclusive []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_policies_exclusive.json
var awsIamUserPoliciesExclusive []byte

//go:embed mapping/aws/resource/m2/aws_m2_application.json
var awsM2Application []byte

//go:embed mapping/aws/resource/m2/aws_m2_deployment.json
var awsM2Deployment []byte

//go:embed mapping/aws/resource/m2/aws_m2_environment.json
var awsM2Environment []byte

//go:embed mapping/aws/resource/memorydb/aws_memorydb_user.json
var awsMemorydbUser []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_human_task_ui.json
var awsSagemakerHumanTaskUI []byte

//go:embed mapping/aws/resource/cloudfront-keyvaluestore/aws_cloudfrontkeyvaluestore_key.json
var awsCloudfrontkeyvaluestoreKey []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_tag.json
var awsEcsTag []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_trust_store.json
var awsLbTrustStore []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_trust_store_revocation.json
var awsLbTrustStoreRevocation []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_folder.json
var awsQuicksightFolder []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_group.json
var awsQuicksightGroup []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_group_membership.json
var awsQuicksightGroupMembership []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_namespace.json
var awsQuicksightNamespace []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_user.json
var awsQuicksightUser []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_asset_type.json
var awsDatazoneAssetType []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_environment.json
var awsDatazoneEnvironment []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_environment_profile.json
var awsDatazoneEnvironmentProfile []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_form_type.json
var awsDatazoneFormType []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_glossary.json
var awsDatazoneGlossary []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_glossary_term.json
var awsDatazoneGlossaryTerm []byte

//go:embed mapping/aws/resource/datazone/aws_datazone_user_profile.json
var awsDatazoneUserProfile []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_account_subscription.json
var awsQuicksightAccountSubscription []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_folder_membership.json
var awsQuicksightFolderMembership []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_iam_policy_assignment.json
var awsQuicksightIamPolicyAssignment []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_ingestion.json
var awsQuicksightIngestion []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_template_alias.json
var awsQuicksightTemplateAlias []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_vpc_connection.json
var awsQuicksightVpcConnection []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_analytics_configuration.json
var awsS3BucketAnalyticsConfiguration []byte

//go:embed mapping/aws/resource/backup/aws_backup_logically_air_gapped_vault.json
var awsBackupLogicallyAirGappedVault []byte

//go:embed mapping/aws/resource/kinesis/aws_kinesis_resource_policy.json
var awsKinesisResourcePolicy []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_deployment.json
var awsAppconfigDeployment []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_graphql_api.json
var awsAppsyncGraphAPI []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrock_inference_profile.json
var awsBedrockInferenceProfile []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_dx_gateway_attachment.json
var awsNetworkmanagerDxGatewayAttachment []byte

//go:embed mapping/aws/resource/s3tables/aws_s3tables_namespace.json
var awsS3tablesNamespace []byte

//go:embed mapping/aws/resource/s3tables/aws_s3tables_table.json
var awsS3tablesTable []byte

//go:embed mapping/aws/resource/s3tables/aws_s3tables_table_bucket_policy.json
var awsS3tablesTableBucketPolicy []byte

//go:embed mapping/aws/resource/s3tables/aws_s3tables_table_policy.json
var awsS3tablesTablePolicy []byte

//go:embed mapping/aws/resource/s3tables/aws_s3tables_table_bucket.json
var awsS3tablesTableBucket []byte

//go:embed mapping/aws/resource/apigateway/aws_api_gateway_domain_name_access_association.json
var awsAPIGatewayDomainNameAccessAssociation []byte

//go:embed mapping/aws/resource/appconfig/aws_appconfig_deployment_strategy.json
var awsAppconfigDeploymentStrategy []byte

//go:embed mapping/aws/resource/appmesh/aws_appmesh_gateway_route.json
var awsAppmeshGatewayRoute []byte

//go:embed mapping/aws/resource/appmesh/aws_appmesh_mesh.json
var awsAppmeshMesh []byte

//go:embed mapping/aws/resource/appmesh/aws_appmesh_route.json
var awsAppmeshRoute []byte

//go:embed mapping/aws/resource/appmesh/aws_appmesh_virtual_gateway.json
var awsAppmeshVirtualGateway []byte

//go:embed mapping/aws/resource/appmesh/aws_appmesh_virtual_node.json
var awsAppmeshVirtualNode []byte

//go:embed mapping/aws/resource/appmesh/aws_appmesh_virtual_router.json
var awsAppmeshVirtualRouter []byte

//go:embed mapping/aws/resource/appmesh/aws_appmesh_virtual_service.json
var awsAppmeshVirtualService []byte

//go:embed mapping/aws/resource/appstream/aws_appstream_directory_config.json
var awsAppstreamDirectoryConfig []byte

//go:embed mapping/aws/resource/appstream/aws_appstream_fleet.json
var awsAppstreamFleet []byte

//go:embed mapping/aws/resource/appstream/aws_appstream_fleet_stack_association.json
var awsAppstreamFleetStackAssociation []byte

//go:embed mapping/aws/resource/appstream/aws_appstream_stack.json
var awsAppstreamStack []byte

//go:embed mapping/aws/resource/appstream/aws_appstream_user.json
var awsAppstreamUser []byte

//go:embed mapping/aws/resource/appstream/aws_appstream_user_stack_association.json
var awsAppstreamUserStackAssociation []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_api_cache.json
var awsAppsyncAPICache []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_api_key.json
var awsAppsyncAPIKey []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_datasource.json
var awsAppsyncDatasource []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_source_api_association.json
var awsAppsyncSourceAPIAssociation []byte

//go:embed mapping/aws/resource/appsync/aws_appsync_type.json
var awsAppsyncType []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrock_guardrail.json
var awsBedrockGuardrail []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrock_guardrail_version.json
var awsBedrockGuardrailVersion []byte

//go:embed mapping/aws/resource/cleanrooms/aws_cleanrooms_membership.json
var awsCleanroomsMembership []byte

//go:embed mapping/aws/resource/codeconnections/aws_codeconnections_connection.json
var awsCodeconnectionsConnection []byte

//go:embed mapping/aws/resource/profile/aws_customerprofiles_domain.json
var awsCustomerprofilesDomain []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_instance_fleet.json
var awsEmrInstanceFleet []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emr_instance_group.json
var awsEmrInstanceGroup []byte

//go:embed mapping/aws/resource/globalaccelerator/aws_globalaccelerator_custom_routing_accelerator.json
var awsGlobalacceleratorCustomRoutingAccelerator []byte

//go:embed mapping/aws/resource/globalaccelerator/aws_globalaccelerator_custom_routing_endpoint_group.json
var awsGlobalacceleratorCustomRoutingEndpointGroup []byte

//go:embed mapping/aws/resource/globalaccelerator/aws_globalaccelerator_custom_routing_listener.json
var awsGlobalacceleratorCustomRoutingListener []byte

//go:embed mapping/aws/resource/glue/aws_glue_partition.json
var awsGluePartition []byte

//go:embed mapping/aws/resource/identitystore/aws_identitystore_group.json
var awsIdentitystoreGroup []byte

//go:embed mapping/aws/resource/identitystore/aws_identitystore_user.json
var awsIdentitystoreUser []byte

//go:embed mapping/aws/resource/identitystore/aws_identitystore_group_membership.json
var awsIdentitystoreGroupMembership []byte

//go:embed mapping/aws/resource/imagebuilder/aws_imagebuilder_lifecycle_policy.json
var awsImagebuilderLifecyclePolicy []byte

//go:embed mapping/aws/resource/iot/aws_iot_domain_configuration.json
var awsIotDomainConfiguration []byte

//go:embed mapping/aws/resource/codebuild/aws_codebuild_fleet.json
var awsCodeBuildFleet []byte

//go:embed mapping/aws/resource/config/aws_config_aggregate_authorization.json
var awsConfigAggregateAuthorization []byte

//go:embed mapping/aws/resource/config/aws_config_organization_managed_rule.json
var awsConfigOrganizationManagedRule []byte

//go:embed mapping/aws/resource/config/aws_config_remediation_configuration.json
var awsConfigRemediationConfiguration []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_instance_connect_endpoint.json
var awsEc2InstanceConnectEndpoint []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_traffic_mirror_filter_rule.json
var awsEc2TrafficMirrorFilterRule []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_traffic_mirror_session.json
var awsEc2TrafficMirrorSession []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_traffic_mirror_filter.json
var awsEc2TrafficMirrorFilter []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_traffic_mirror_target.json
var awsEc2TrafficMirrorTarget []byte

//go:embed mapping/aws/resource/glue/aws_glue_data_quality_ruleset.json
var awsGlueDataQualityRuleset []byte

//go:embed mapping/aws/resource/glue/aws_glue_dev_endpoint.json
var awsGlueDevEndpoint []byte

//go:embed mapping/aws/resource/grafana/aws_grafana_workspace.json
var awsGrafanaWorkspace []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_data_lake_settings.json
var awsLakeformationDataLakeSettings []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_permissions.json
var awsLakeformationPermissions []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_resource.json
var awsLakeformationResource []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_delivery.json
var awsCloudwatchLogDelivery []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_delivery_destination.json
var awsCloudwatchLogDeliveryDestination []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_delivery_destination_policy.json
var awsCloudwatchLogDeliveryDestinationPolicy []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_delivery_source.json
var awsCloudwatchLogDeliverySource []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_index_policy.json
var awsCloudwatchLogIndexPolicy []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_anomaly_detector.json
var awsCloudwatchLogAnomalyDetector []byte

//go:embed mapping/aws/resource/route53profiles/aws_route53profiles_association.json
var awsRoute53profilesAssociation []byte

//go:embed mapping/aws/resource/route53profiles/aws_route53profiles_profile.json
var awsRoute53profilesProfile []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalogappregistry_attribute_group.json
var awsServicecatalogappregistryAttributeGroup []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalogappregistry_attribute_group_association.json
var awsServicecatalogappregistryAttributeGroupAssociation []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_server.json
var awsTransferServer []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_ssh_key.json
var awsTransferSSHKey []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_user.json
var awsTransferUser []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_block_public_access_exclusion.json
var awsVpcBlockPublicAccessExclusion []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_block_public_access_options.json
var awsVpcBlockPublicAccessOptions []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_resource_configuration.json
var awsVpclatticeResourceConfiguration []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_resource_gateway.json
var awsVpclatticeResourceGateway []byte

//go:embed mapping/aws/resource/amplify/aws_amplify_backend_environment.json
var awsAmplifyBackendEnvironment []byte

//go:embed mapping/aws/resource/amplify/aws_amplify_webhook.json
var awsAmplifyWebhook []byte

//go:embed mapping/aws/resource/appfabric/aws_appfabric_app_authorization.json
var awsAppfabricAppAuthorization []byte

//go:embed mapping/aws/resource/appfabric/aws_appfabric_app_authorization_connection.json
var awsAppfabricAppAuthorizationConnection []byte

//go:embed mapping/aws/resource/appfabric/aws_appfabric_app_bundle.json
var awsAppfabricAppBundle []byte

//go:embed mapping/aws/resource/appfabric/aws_appfabric_ingestion.json
var awsAppfabricIngestion []byte

//go:embed mapping/aws/resource/appfabric/aws_appfabric_ingestion_destination.json
var awsAppfabricIngestionDestination []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_connection.json
var awsApprunnerConnection []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_custom_domain_association.json
var awsApprunnerCustomDomainAssociation []byte

//go:embed mapping/aws/resource/apprunner/aws_apprunner_deployment.json
var awsApprunnerDeployment []byte

//go:embed mapping/aws/resource/cloud9/aws_cloud9_environment_membership.json
var awsCloud9EnvironmentMembership []byte

//go:embed mapping/aws/resource/cloudformation/aws_cloudcontrolapi_resource.json
var awsCloudcontrolapiResource []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_flow_definition.json
var awsSagemakerFlowDefinition []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_hub.json
var awsSagemakerHub []byte

//go:embed mapping/aws/resource/sagemaker/aws_sagemaker_mlflow_tracking_server.json
var awsSagemakerMlflowTrackingServer []byte

//go:embed mapping/aws/resource/shield/aws_shield_subscription.json
var awsShieldSubscription []byte

//go:embed mapping/aws/resource/ec2/aws_spot_datafeed_subscription.json
var awsSpotDatafeedSubscription []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_activation.json
var awsSsmActivation []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_default_patch_baseline.json
var awsSsmDefaultPatchBaseline []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_service_setting.json
var awsSsmServiceSetting []byte

//go:embed mapping/aws/resource/ssm-quicksetup/aws_ssmquicksetup_configuration_manager.json
var awsSsmquicksetupConfigurationManager []byte

//go:embed mapping/aws/resource/wafregional/aws_wafregional_web_acl_association.json
var awsWafregionalWebACLAssociation []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_agent_collaborator.json
var awsBedrockagentAgentCollaborator []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_agent_knowledge_base_association.json
var awsBedrockagentKnowledgeBaseAssociation []byte

//go:embed mapping/aws/resource/cloudformation/aws_cloudformation_stack_instances.json
var awsCloudformationStackInstances []byte

//go:embed mapping/aws/resource/cloudfront/aws_cloudfront_vpc_origin.json
var awsCloudfrontVpcOrigin []byte

//go:embed mapping/aws/resource/cloudhsm/aws_cloudhsm_v2_cluster.json
var awsCloudhsmV2Vluster []byte

//go:embed mapping/aws/resource/cloudhsm/aws_cloudhsm_v2_hsm.json
var awsCloudhsmV2Hsm []byte

//go:embed mapping/aws/resource/cloudtrail/aws_cloudtrail_organization_delegated_admin_account.json
var awsCloudtrailOrganizationDelegatedAdminAccount []byte

//go:embed mapping/aws/resource/codeconnections/aws_codeconnections_host.json
var awsCodeconnectionsHost []byte

//go:embed mapping/aws/resource/comprehend/aws_comprehend_entity_recognizer.json
var awsComprehendEntityRecognizer []byte

//go:embed mapping/aws/resource/compute-optimizer/aws_computeoptimizer_enrollment_status.json
var awsComputeoptimizerEnrollmentStatus []byte

//go:embed mapping/aws/resource/compute-optimizer/aws_computeoptimizer_recommendation_preferences.json
var awsComputeoptimizerRecommendationPreferences []byte

//go:embed mapping/aws/resource/config/aws_config_organization_custom_policy_rule.json
var awsConfigOrganizationCustomPolicyRule []byte

//go:embed mapping/aws/resource/config/aws_config_organization_custom_rule.json
var awsConfigOrganizationCustomRule []byte

//go:embed mapping/aws/resource/config/aws_config_retention_configuration.json
var awsConfigRetentionConfiguration []byte

//go:embed mapping/aws/resource/cost-optimization-hub/aws_costoptimizationhub_enrollment_status.json
var awsCostoptimizationhubEnrollmentStatus []byte

//go:embed mapping/aws/resource/cost-optimization-hub/aws_costoptimizationhub_preferences.json
var awsCostoptimizationhubPreferences []byte

//go:embed mapping/aws/resource/profile/aws_customerprofiles_profile.json
var awsCustomerprofilesProfile []byte

//go:embed mapping/aws/resource/connect/aws_connect_lambda_function_association.json
var awsConnectLambdaFunctionAssociation []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_account.json
var awsSecurityhubAccount []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_action_target.json
var awsSecurityhubActionTarget []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_configuration_policy_association.json
var awsSecurityhubConfigurationPolicyAssociation []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_organization_admin_account.json
var awsSecurityhubOrganizationAdminAccount []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_standards_control_association.json
var awsSecurityhubStandardsControlAssociation []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_standards_subscription.json
var awsSecurityhubStandardsSubscription []byte

//go:embed mapping/aws/resource/securitylake/aws_securitylake_custom_log_source.json
var awsSecuritylakeCustomLogSource []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_access.json
var awsTransferAccess []byte

//go:embed mapping/aws/resource/transfer/aws_transfer_tag.json
var awsTransferTag []byte

//go:embed mapping/aws/resource/detective/aws_detective_member.json
var awsDetectiveMember []byte

//go:embed mapping/aws/resource/detective/aws_detective_organization_configuration.json
var awsDetectiveOrganizationConfiguration []byte

//go:embed mapping/aws/resource/eks/aws_eks_access_policy_association.json
var awsEksAccessPolicyAssociation []byte

//go:embed mapping/aws/resource/elasticmapreduce/aws_emrcontainers_job_template.json
var awsEmrcontainersJobTemplate []byte

//go:embed mapping/aws/resource/fms/aws_fms_admin_account.json
var awsFmsAdminAccount []byte

//go:embed mapping/aws/resource/glue/aws_glue_catalog_table_optimizer.json
var awsGlueCatalogTableOptimizer []byte

//go:embed mapping/aws/resource/glue/aws_glue_partition_index.json
var awsGluePartitionIndex []byte

//go:embed mapping/aws/resource/grafana/aws_grafana_license_association.json
var awsGrafanaLicenseAssociation []byte

//go:embed mapping/aws/resource/grafana/aws_grafana_workspace_service_account.json
var awsGrafanaWorkspaceServiceAccount []byte

//go:embed mapping/aws/resource/grafana/aws_grafana_workspace_service_account_token.json
var awsGrafanaWorkspaceServiceAccountToken []byte

//go:embed mapping/aws/resource/ec2/aws_internet_gateway_attachment.json
var awsInternetGatewayAttachment []byte

//go:embed mapping/aws/resource/kendra/aws_kendra_experience.json
var awsKendraExperience []byte

//go:embed mapping/aws/resource/kendra/aws_kendra_query_suggestions_block_list.json
var awsKendraQuerySuggestionsBlockList []byte

//go:embed mapping/aws/resource/kendra/aws_kendra_thesaurus.json
var awsKendraThesaurus []byte

//go:embed mapping/aws/resource/aps/aws_prometheus_alert_manager_definition.json
var awsPrometheusAlertManagerDefinition []byte

//go:embed mapping/aws/resource/aps/aws_prometheus_scraper.json
var awsPrometheusScraper []byte

//go:embed mapping/aws/resource/aps/aws_prometheus_workspace.json
var awsPrometheusWorkspace []byte

//go:embed mapping/aws/resource/rds/aws_rds_certificate.json
var awsRdsCertificate []byte

//go:embed mapping/aws/resource/rds/aws_rds_cluster_snapshot_copy.json
var awsRdsClusterSnapshotCopy []byte

//go:embed mapping/aws/resource/rds/aws_rds_custom_db_engine_version.json
var awsRdsCustomDBEngineVersion []byte

//go:embed mapping/aws/resource/rds/aws_rds_export_task.json
var awsRdsExportTask []byte

//go:embed mapping/aws/resource/rds/aws_rds_reserved_instance.json
var awsRdsReservedInstance []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_opt_in.json
var awsLakeformationOptIn []byte
