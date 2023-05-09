package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/aws/data/cognito-idp/aws_cognito_user_pools.json
var dataAwsCognitoUserPools []byte

//go:embed mapping/aws/data/cognito-idp/aws_cognito_user_pool_client.json
var dataAwsCognitoUserPoolClient []byte

//go:embed mapping/aws/data/cognito-idp/aws_cognito_user_pool_clients.json
var dataAwsCognitoUserPoolClients []byte

//go:embed mapping/aws/data/cognito-idp/aws_cognito_user_pool_signing_certificate.json
var dataAwsCognitoUserPoolSigningCertificate []byte

//go:embed mapping/aws/data/template.json
var placeholder []byte

//go:embed mapping/aws/data/ec2/aws_vpcs.json
var dataAwsVpcs []byte

//go:embed mapping/aws/data/ec2/aws_subnets_ids.json
var dataAwsSubnetIds []byte

//go:embed mapping/aws/data/ec2/aws_ami.json
var dataAwsAmi []byte

//go:embed mapping/aws/data/ec2/aws_availability_zones.json
var dataAwsAvailabilityZones []byte

//go:embed mapping/aws/data/iam/aws_iam_policy.json
var dataAwsIamPolicy []byte

//go:embed mapping/aws/data/iam/aws_iam_role.json
var dataAwsIamRole []byte

//go:embed mapping/aws/data/ec2/aws_vpc.json
var dataAwsVpc []byte

//go:embed mapping/aws/data/s3/aws_s3_bucket.json
var dataAwsS3Bucket []byte

//go:embed mapping/aws/data/inspector/aws_inspector_rules_packages.json
var dataAwsInspectorRulesPackages []byte

//go:embed mapping/aws/data/route53/aws_route53_zone.json
var dataAwsRoute53Zone []byte

//go:embed mapping/aws/data/ec2/aws_security_group.json
var dataAwsSecurityGroup []byte

//go:embed mapping/aws/data/sns/aws_sns_topic.json
var dataAwsSnsTopic []byte

//go:embed mapping/aws/data/ssm/aws_ssm_parameter.json
var dataAwsSsmParameter []byte

//go:embed mapping/aws/data/kms/aws_kms_alias.json
var dataAwsKmsAlias []byte

//go:embed mapping/aws/data/kms/aws_kms_ciphertext.json
var dataAwsKmsCiphertext []byte

//go:embed mapping/aws/data/kms/aws_kms_key.json
var dataAwsKmsKey []byte

//go:embed mapping/aws/data/route53/aws_route_tables.json
var dataAwsRouteTables []byte

//go:embed mapping/aws/data/elasticbeanstalk/aws_elastic_beanstalk_solution_stack.json
var dataAwsElasticBeanstalkSolutionStack []byte

//go:embed mapping/aws/data/organizations/aws_organizations_organization.json
var dataAwsOrganizationsOrganization []byte

//go:embed mapping/aws/data/ec2/aws_ebs_default_kms_key.json
var dataAwsEbsDefaultKmsKey []byte

//go:embed mapping/aws/data/wafv2/aws_wafv2_ip_set.json
var dataAwsWafv2IpSet []byte

//go:embed mapping/aws/data/wafv2/aws_wafv2_web_acl.json
var dataAwsWafv2WebACL []byte

//go:embed mapping/aws/data/wafv2/aws_wafv2_rule_group.json
var dataAwsWafv2RuleGroup []byte

//go:embed mapping/aws/data/wafv2/aws_wafv2_regex_pattern_set.json
var dataAwsWafv2RegexPatternSet []byte

//go:embed mapping/aws/data/cloudwatch/aws_cloudwatch_log_group.json
var dataAwsCloudwatchLogGroup []byte

//go:embed mapping/aws/data/cloudwatch/aws_cloudwatch_log_groups.json
var dataAwsCloudwatchLogGroups []byte

//go:embed mapping/aws/data/directoryservice/aws_directory_service_directory.json
var dataAwsDirectoryServiceDirectory []byte

//go:embed mapping/aws/data/sso/aws_ssoadmin_instances.json
var dataAwsSsoadminInstances []byte

//go:embed mapping/aws/data/ecr/aws_ecr_authorization_token.json
var dataAwsEcrAuthorization []byte

//go:embed mapping/aws/data/ecs/aws_ecs_cluster.json
var dataAwsEcsCluster []byte

//go:embed mapping/aws/data/ecs/aws_ecs_task_definition.json
var dataAwsEcsTaskDefinition []byte

//go:embed mapping/aws/data/eks/aws_eks_cluster.json
var dataAwsEksCluster []byte

//go:embed mapping/aws/data/lambda/aws_lambda_function.json
var dataAwsLambdaFunction []byte

//go:embed mapping/aws/data/ecr/aws_ecr_authorization_token.json
var dataAwsEcrAuthorizationToken []byte

//go:embed mapping/aws/data/ecr/aws_ecr_repository.json
var dataAwsEcrRepository []byte

//go:embed mapping/aws/data/outposts/aws_outposts_outpost.json
var dataAwsOutpostsOutpost []byte

//go:embed mapping/aws/data/rds/aws_rds_engine_version.json
var dataAwsRdsEngineVersion []byte

//go:embed mapping/aws/data/rds/aws_db_cluster_snapshot.json
var dataAwsDbClusterSnapshot []byte

//go:embed mapping/aws/data/rds/aws_db_event_categories.json
var dataAwsDbEventCategories []byte

//go:embed mapping/aws/data/rds/aws_db_instance.json
var dataAwsDbInstance []byte

//go:embed mapping/aws/data/rds/aws_db_snapshot.json
var dataAwsDbSnapshot []byte

//go:embed mapping/aws/data/rds/aws_db_subnet_group.json
var dataAwsDbSubnetGroup []byte

//go:embed mapping/aws/data/rds/aws_rds_certificate.json
var dataAwsRdsCertificate []byte

//go:embed mapping/aws/data/rds/aws_rds_cluster.json
var dataAwsRdsCluster []byte

//go:embed mapping/aws/data/rds/aws_rds_orderable_db_instance.json
var dataAwsRdsOrderableDbInstance []byte

//go:embed mapping/aws/data/ec2/aws_vpc_endpoint_service.json
var dataAwsVpcEndpointService []byte

//go:embed mapping/aws/data/redshift/aws_redshift_cluster.json
var dataAwsRedshiftCluster []byte

//go:embed mapping/aws/data/redshift/aws_redshift_orderable_cluster.json
var dataAwsRedshiftOrderableCluster []byte

//go:embed mapping/aws/data/redshift/aws_redshift_cluster_credentials.json
var dataAwsRedshiftClusterCredentials []byte

//go:embed mapping/aws/data/redshift/aws_redshift_subnet_group.json
var dataAwsRedshiftSubnetGroup []byte

//go:embed mapping/aws/data/glue/aws_glue_connection.json
var dataAwsGlueConnection []byte

//go:embed mapping/aws/data/glue/aws_glue_script.json
var dataAwsGlueScript []byte

//go:embed mapping/aws/data/glue/aws_glue_data_catalog_encryption_settings.json
var dataAwsDataCatalogEncryptionSettings []byte

//go:embed mapping/aws/data/elasticfilesystem/aws_efs_mount_target.json
var dataAwsEfsMountTarget []byte

//go:embed mapping/aws/data/elasticfilesystem/aws_efs_file_system.json
var dataAwsEfsFileSystem []byte

//go:embed mapping/aws/data/elasticfilesystem/aws_efs_access_point.json
var dataAwsEfsAccessPoint []byte

//go:embed mapping/aws/data/elasticfilesystem/aws_efs_access_points.json
var dataAwsEfsAccessPoints []byte

//go:embed mapping/aws/data/iam/aws_iam_account_alias.json
var dataAwsIamAccountAlias []byte

//go:embed mapping/aws/data/iam/aws_iam_group.json
var dataAwsIamGroup []byte

//go:embed mapping/aws/data/iam/aws_iam_instance_profile.json
var dataAwsIamInstanceProfile []byte

//go:embed mapping/aws/data/iam/aws_iam_instance_profiles.json
var dataAwsIamInstanceProfiles []byte

//go:embed mapping/aws/data/iam/aws_iam_openid_connect_provider.json
var dataAwsIamOpenIDConnectProvider []byte

//go:embed mapping/aws/data/iam/aws_iam_saml_provider.json
var dataAwsIamSamlProvider []byte

//go:embed mapping/aws/data/iam/aws_iam_server_certificate.json
var dataAwsIamServerCertificate []byte

//go:embed mapping/aws/data/iam/aws_iam_user.json
var dataAwsIamUser []byte

//go:embed mapping/aws/data/iam/aws_iam_user_ssh_key.json
var dataAwsIamUserSSHKey []byte

//go:embed mapping/aws/data/iam/aws_iam_users.json
var dataAwsIamUsers []byte

//go:embed mapping/aws/data/ec2/aws_ec2_managed_prefix_list.json
var dataAwsEc2ManagedPrefixList []byte

//go:embed mapping/aws/data/ec2/aws_ec2_network_insights_analysis.json
var dataAwsEc2NetworkInsightsAnalysis []byte

//go:embed mapping/aws/data/ec2/aws_ec2_network_insights_path.json
var dataAwsEc2NetworkInsightsPath []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_attachment.json
var dataAwsEc2TransitGatewayAttachment []byte

//go:embed mapping/aws/data/geo/aws_location_tracker_association.json
var dataAwsLocationTrackerAssociation []byte

//go:embed mapping/aws/data/geo/aws_location_tracker_associations.json
var dataAwsLocationTrackerAssociations []byte

//go:embed mapping/aws/data/workspaces/aws_workspaces_bundle.json
var dataAwsWorkspacesBundle []byte

//go:embed mapping/aws/data/ec2/aws_route_table.json
var dataAwsRouteTable []byte

//go:embed mapping/aws/data/appconfig/aws_appconfig_configuration_profiles.json
var dataAwsAppconfigConfigurationProfiles []byte

//go:embed mapping/aws/data/appconfig/aws_appconfig_environment.json
var dataAwsAppconfigEnvironment []byte

//go:embed mapping/aws/data/appconfig/aws_appconfig_environments.json
var dataAwsAppconfigEnvironments []byte

//go:embed mapping/aws/data/kms/aws_kms_custom_key_store.json
var dataAwsKmsCustomKeyStore []byte

//go:embed mapping/aws/data/ec2/aws_vpc_ipam_pool_cidrs.json
var dataAwsVpcIpamPoolCidrs []byte

//go:embed mapping/aws/data/elasticbeanstalk/aws_elastic_beanstalk_application.json
var dataAwsElasticBeanstalkApplication []byte

//go:embed mapping/aws/data/cloudwatch/aws_cloudwatch_event_bus.json
var dataCloudwatchEventBus []byte

//go:embed mapping/aws/data/cloudwatch/aws_cloudwatch_event_connection.json
var dataCloudwatchEventConnection []byte

//go:embed mapping/aws/data/cloudwatch/aws_cloudwatch_event_source.json
var dataCloudwatchEventSource []byte

//go:embed mapping/aws/data/backup/aws_backup_framework.json
var dataBackupFramework []byte

//go:embed mapping/aws/data/backup/aws_backup_plan.json
var dataBackupPlan []byte

//go:embed mapping/aws/data/backup/aws_backup_report_plan.json
var dataBackupReportPlan []byte

//go:embed mapping/aws/data/backup/aws_backup_selection.json
var dataBackupSelection []byte

//go:embed mapping/aws/data/ec2/aws_ebs_snapshot.json
var dataAwsEbsSnapshot []byte

//go:embed mapping/aws/data/ec2/aws_ebs_snapshot_ids.json
var dataAwsEbsSnapshotIds []byte

//go:embed mapping/aws/data/ec2/aws_ebs_volume.json
var dataAwsEbsVolume []byte

//go:embed mapping/aws/data/ec2/aws_ebs_volumes.json
var dataAwsEbsVolumes []byte

//go:embed mapping/aws/data/apigateway/aws_api_gateway.json
var dataAwsAPIGateway []byte

//go:embed mapping/aws/data/servicequota/aws_servicequotas_service.json
var dataAwsServicequotasService []byte

//go:embed mapping/aws/data/servicequota/aws_servicequotas_service_quota.json
var dataAwsServicequotaServiceQuota []byte

//go:embed mapping/aws/data/kafka/aws_msk_broker_nodes.json
var dataAwsBrokerNodes []byte

//go:embed mapping/aws/data/kafka/aws_msk_cluster.json
var dataAwsMskCluster []byte

//go:embed mapping/aws/data/kafka/aws_msk_configuration.json
var dataAwsMskConfiguration []byte

//go:embed mapping/aws/data/kafka/aws_msk_kafka_version.json
var dataAwsMskKafkaVersion []byte

//go:embed mapping/aws/data/acm/aws_acm_certificate.json
var dataAwsAcmCertificate []byte

//go:embed mapping/aws/data/autoscaling/aws_autoscaling_group.json
var dataAwsAutoscalingGroup []byte

//go:embed mapping/aws/data/autoscaling/aws_autoscaling_groups.json
var dataAwsAutoscalingGroups []byte

//go:embed mapping/aws/data/appconfig/aws_appconfig_configuration_profile.json
var dataAwsAppconfigConfigurationProfile []byte

//go:embed mapping/aws/data/appmesh/aws_appmesh_mesh.json
var dataAwsAppmeshMesh []byte

//go:embed mapping/aws/data/appmesh/aws_appmesh_virtual_service.json
var dataAppmeshVirtualService []byte

//go:embed mapping/aws/data/backup/aws_backup_vault.json
var dataAwsBackupVault []byte

//go:embed mapping/aws/data/autoscaling/aws_launch_configuration.json
var dataAwsLaunchConfiguration []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway.json
var dataAwsEc2Transitgateway []byte

//go:embed mapping/aws/data/elasticloadbalancing/aws_lb.json
var dataAwsLb []byte

//go:embed mapping/aws/data/secretsmanager/aws_secretsmanager_secret.json
var dataAwsSecretsmanagerSecret []byte

//go:embed mapping/aws/data/secretsmanager/aws_secretsmanager_secret_version.json
var dataAwsSecretsmanagerSecretVersion []byte

//go:embed mapping/aws/data/ses/aws_sesv2_dedicated_ip_pool.json
var dataAwsSesv2DedicatedIPPool []byte

//go:embed mapping/aws/data/sqs/aws_sqs_queues.json
var dataAwsSqsQueues []byte

//go:embed mapping/aws/data/ec2/aws_vpc_ipam_pools.json
var dataAwsVpcIpamPools []byte

//go:embed mapping/aws/data/auditmanager/aws_auditmanager_control.json
var dataAwsAuditmanagerControl []byte

//go:embed mapping/aws/data/auditmanager/aws_auditmanager_framework.json
var dataAwsAuditmanagerFramework []byte

//go:embed mapping/aws/data/connect/aws_connect_instance_storage_config.json
var dataAwsConnectInstanceStorageConfig []byte

//go:embed mapping/aws/data/controltower/aws_controltower_controls.json
var dataAwsControltowerControls []byte

//go:embed mapping/aws/data/rds/aws_db_instances.json
var dataAwsDbInstances []byte

//go:embed mapping/aws/data/directconnect/aws_dx_router_configuration.json
var dataAwsDxRouterConfiguration []byte

//go:embed mapping/aws/data/dynamodb/aws_dynamodb_table_item.json
var dataAwsDynamodbTableItem []byte

//go:embed mapping/aws/data/elasticache/aws_elasticache_subnet_group.json
var dataAwsElasticacheSubnetGroup []byte

//go:embed mapping/aws/data/glue/aws_glue_catalog_table.json
var dataAwsGlueCatalogTable []byte

//go:embed mapping/aws/data/ivs/aws_ivs_stream_key.json
var dataAwsIvsStreamKey []byte

//go:embed mapping/aws/data/elasticloadbalancing/aws_lbs.json
var dataAwsLbs []byte

//go:embed mapping/aws/data/rds/aws_rds_reserved_instance_offering.json
var dataAwsRdsReservedInstanceOffering []byte

//go:embed mapping/aws/data/route53/aws_route53_resolver_firewall_rules.json
var dataAwsRoute53ResolverFirewallConfig []byte

//go:embed mapping/aws/data/route53/aws_route53_resolver_firewall_domain_list.json
var dataAwsRoute53ResolverFirewallDomainList []byte

//go:embed mapping/aws/data/route53/aws_route53_resolver_firewall_rule_group.json
var dataAwsRoute53ResolverFirewallRuleGroup []byte

//go:embed mapping/aws/data/route53/aws_route53_resolver_firewall_rule_group_association.json
var dataAwsRoute53ResolverFirewallGroupAssociation []byte

//go:embed mapping/aws/data/route53/aws_route53_resolver_firewall_rules.json
var dataAwsRoute53ResolverFirewallRules []byte

//go:embed mapping/aws/data/s3/aws_s3control_multi_region_access_point.json
var dataAwsS3controlMultiRegionAccessPoint []byte

//go:embed mapping/aws/data/dynamodb/aws_dynamodb_table.json
var dataAwsDynamodbTable []byte

//go:embed mapping/aws/data/codeartifact/aws_codeartifact_authorization_token.json
var datAwsCodeartifactAutorization []byte

//go:embed mapping/aws/data/codeartifact/aws_codeartifact_repository_endpoint.json
var dataAwsCodeartifactRepositoryEndpoint []byte

//go:embed mapping/aws/data/iam/aws_iam_roles.json
var dataAwsIamRoles []byte

//go:embed mapping/aws/data/cloudformation/aws_cloudformation_export.json
var dataAwsCloudformationExport []byte

//go:embed mapping/aws/data/cloudformation/aws_cloudformation_stack.json
var dataAwsCloudformationStack []byte

//go:embed mapping/aws/data/cloudformation/aws_cloudformation_type.json
var dataAwsCloudformationType []byte

//go:embed mapping/aws/data/cloudhsm/aws_cloudhsm_v2_cluster.json
var dataAwsCloudhsmV2Cluster []byte

//go:embed mapping/aws/data/codecommit/aws_codecommit_approval_rule_template.json
var dataAwsCodecommitApprovalRuleTemplate []byte

//go:embed mapping/aws/data/codecommit/aws_codecommit_repository.json
var dataAwsCodecommitRepository []byte

//go:embed mapping/aws/data/codestar-connections/aws_codestarconnections_connection.json
var dataAwsCodestarconnectionsConnection []byte

//go:embed mapping/aws/data/batch/aws_batch_compute_environment.json
var dataAwsBatchComputeEnvironment []byte

//go:embed mapping/aws/data/batch/aws_batch_job_queue.json
var dataAwsBatchJobQueue []byte

//go:embed mapping/aws/data/batch/aws_batch_scheduling_policy.json
var dataAwsBatchSchedulingPolicy []byte

//go:embed mapping/aws/data/ce/aws_ce_cost_category.json
var dataAwsCeCostCategory []byte

//go:embed mapping/aws/data/ce/aws_ce_tags.json
var dataAwsCeTags []byte

//go:embed mapping/aws/data/cloudformation/aws_cloudcontrolapi_resource.json
var dataAwsCloudcontrolapiResource []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_cache_policy.json
var dataAwsCloudfrontCachePolicy []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_distribution.json
var dataAwsCloudfrontDistribution []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_function.json
var dataAwsCloudfrontFunction []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_origin_access_identities.json
var dataAwsCloudfrontOriginAccessIdentities []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_origin_access_identity.json
var dataAwsCloudfrontOriginAccessIdentity []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_origin_request_policy.json
var dataAwsCloudfrontOriginRequestPolicy []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_realtime_log_config.json
var dataAwsCloudfrontRealtimeLogConfig []byte

//go:embed mapping/aws/data/cloudfront/aws_cloudfront_response_headers_policy.json
var dataAwsCloudfrontResponseHeadersPolicy []byte

//go:embed mapping/aws/data/servicediscovery/aws_service_discovery_dns_namespace.json
var dataAwsServiceDiscoveryDNSNamespace []byte

//go:embed mapping/aws/data/servicediscovery/aws_service_discovery_http_namespace.json
var dataAwsServiceDiscoveryHTTPNamespace []byte

//go:embed mapping/aws/data/servicediscovery/aws_service_discovery_service.json
var dataAwsServiceDiscoveryService []byte

//go:embed mapping/aws/data/connect/aws_connect_bot_association.json
var dataAwsConnectBotAssociation []byte

//go:embed mapping/aws/data/connect/aws_connect_contact_flow.json
var dataAwsConnectContactFlow []byte

//go:embed mapping/aws/data/connect/aws_connect_contact_flow_module.json
var dataAwsConnectContactFlowModule []byte

//go:embed mapping/aws/data/connect/aws_connect_hours_of_operation.json
var dataAwsConnectHoursOfOperation []byte

//go:embed mapping/aws/data/connect/aws_connect_instance.json
var dataAwsConnectInstance []byte

//go:embed mapping/aws/data/connect/aws_connect_lambda_function_association.json
var dataAwsConnectLambdaFunctionAssociation []byte

//go:embed mapping/aws/data/connect/aws_connect_prompt.json
var dataAwsConnectPrompt []byte

//go:embed mapping/aws/data/connect/aws_connect_queue.json
var dataAwsConnectQueue []byte

//go:embed mapping/aws/data/connect/aws_connect_quick_connect.json
var dataAwsConnectQuickConnect []byte

//go:embed mapping/aws/data/connect/aws_connect_routing_profile.json
var dataAwsConnectRoutingProfile []byte

//go:embed mapping/aws/data/connect/aws_connect_security_profile.json
var dataAwsConnectSecurityProfile []byte

//go:embed mapping/aws/data/connect/aws_connect_user_hierarchy_group.json
var dataAwsConnectUserHierarchyGroup []byte

//go:embed mapping/aws/data/connect/aws_connect_user_hierarchy_structure.json
var dataAwsConnectUserHierarchyStructure []byte

//go:embed mapping/aws/data/datapipeline/aws_datapipeline_pipeline.json
var dataAwsDatapipelinePipeline []byte

//go:embed mapping/aws/data/datapipeline/aws_datapipeline_pipeline_definition.json
var dataAwsDatapipelinePipelineDefinition []byte

//go:embed mapping/aws/data/rds/aws_docdb_engine_version.json
var dataAwsDocDBEngineVersion []byte

//go:embed mapping/aws/data/rds/aws_docdb_orderable_db_instance.json
var dataAwsDocDBOrderableDBInstance []byte

//go:embed mapping/aws/data/directconnect/aws_dx_connection.json
var dataAwsDxConnection []byte

//go:embed mapping/aws/data/directconnect/aws_dx_gateway.json
var dataAwsDxGateway []byte

//go:embed mapping/aws/data/directconnect/aws_dx_location.json
var dataAwsDxLocation []byte

//go:embed mapping/aws/data/directconnect/aws_dx_locations.json
var dataAwsDxLocations []byte

//go:embed mapping/aws/data/elasticloadbalancing/aws_lb_target_group.json
var dataAwsLbTargetGroup []byte

//go:embed mapping/aws/data/elasticloadbalancing/aws_lb_listener.json
var dataAwsLbListener []byte

//go:embed mapping/aws/data/elasticache/aws_elasticache_cluster.json
var dataAwsElasticacheCluster []byte

//go:embed mapping/aws/data/elasticache/aws_elasticache_replication_group.json
var dataAwsElasticacheReplicationGroup []byte

//go:embed mapping/aws/data/elasticache/aws_elasticache_user.json
var dataAwsElasticacheUser []byte

//go:embed mapping/aws/data/elasticsearch/aws_elasticsearch_domain.json
var dataAwsElasticsearchDomain []byte

//go:embed mapping/aws/data/elasticmapreduce/aws_emr_release_labels.json
var dataAwsEmrReleaseLabels []byte

//go:embed mapping/aws/data/ecr/aws_ecr_image.json
var dataAwsEcrImage []byte

//go:embed mapping/aws/data/ecs/aws_ecs_container_definition.json
var dataAwsEcsContainerDefinition []byte

//go:embed mapping/aws/data/ecs/aws_ecs_service.json
var dataDataEcsService []byte

//go:embed mapping/aws/data/eks/aws_eks_clusters.json
var dataAwsEksClusters []byte

//go:embed mapping/aws/data/eks/aws_eks_node_group.json
var dataAwsEksNodeGroup []byte

//go:embed mapping/aws/data/eks/aws_eks_node_groups.json
var dataAwsEksNodeGroups []byte
