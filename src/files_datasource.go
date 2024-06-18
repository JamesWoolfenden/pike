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

//go:embed mapping/aws/data/ec2/aws_route_tables.json
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
var dataAwsDBClusterSnapshot []byte

//go:embed mapping/aws/data/rds/aws_db_event_categories.json
var dataAwsDBEventCategories []byte

//go:embed mapping/aws/data/rds/aws_db_instance.json
var dataAwsDBInstance []byte

//go:embed mapping/aws/data/rds/aws_db_snapshot.json
var dataAwsDBSnapshot []byte

//go:embed mapping/aws/data/rds/aws_db_subnet_group.json
var dataAwsDBSubnetGroup []byte

//go:embed mapping/aws/data/rds/aws_rds_certificate.json
var dataAwsRdsCertificate []byte

//go:embed mapping/aws/data/rds/aws_rds_cluster.json
var dataAwsRdsCluster []byte

//go:embed mapping/aws/data/rds/aws_rds_orderable_db_instance.json
var dataAwsRdsOrderableDBInstance []byte

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

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_attachments.json
var dataAwsEc2TransitGatewayAttachments []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_dx_gateway_attachment.json
var dataAwsEc2TransitGatewayDxGatewayAttachment []byte

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
var dataAwsDBInstances []byte

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

//go:embed mapping/aws/data/ec2/aws_network_acls.json
var dataAwsNetworkAcls []byte

//go:embed mapping/aws/data/ec2/aws_prefix_list.json
var dataAwsPrefixList []byte

//go:embed mapping/aws/data/ec2/aws_vpc_dhcp_options.json
var dataAwsVpcDhcpOptions []byte

//go:embed mapping/aws/data/route53/aws_route53_resolver_rule.json
var dataAwsRoute53ResolverRule []byte

//go:embed mapping/aws/data/acm-pca/aws_acmpca_certificate.json
var dataAwsAcmpcaCertificate []byte

//go:embed mapping/aws/data/acm-pca/aws_acmpca_certificate_authority.json
var dataAwsAcmpcaCertificateAuthority []byte

//go:embed mapping/aws/data/appintegrations/aws_appintegrations_event_integration.json
var dataAwsAppintergrationsEventIntegration []byte

//go:embed mapping/aws/data/prometheus/aws_prometheus_workspace.json
var dataAwsPrometheusWorkspace []byte

//go:embed mapping/aws/data/prometheus/aws_prometheus_workspaces.json
var dataAwsPrometheusWorkspaces []byte

//go:embed mapping/aws/data/ec2/aws_ami_ids.json
var dataAwsAmiIds []byte

//go:embed mapping/aws/data/appmesh/aws_appmesh_gateway_route.json
var dataAwsAppmeshGatewayRoute []byte

//go:embed mapping/aws/data/appmesh/aws_appmesh_route.json
var dataAwsAppmeshRoute []byte

//go:embed mapping/aws/data/appmesh/aws_appmesh_virtual_gateway.json
var dataAwsAppmeshVirtualGateway []byte

//go:embed mapping/aws/data/appmesh/aws_appmesh_virtual_node.json
var dataAwsAppmeshVirtualNode []byte

//go:embed mapping/aws/data/appmesh/aws_appmesh_virtual_router.json
var dataAwsAppmeshVirtualRouter []byte

//go:embed mapping/aws/data/ec2/aws_availability_zone.json
var dataAwsAvailabilityZone []byte

//go:embed mapping/aws/data/budgets/aws_budgets_budget.json
var dataAwsBudgetsBudget []byte

//go:embed mapping/aws/data/connect/aws_connect_user.json
var dataAwsConnectUser []byte

//go:embed mapping/aws/data/connect/aws_connect_vocabulary.json
var dataAwsConnectVocabulary []byte

//go:embed mapping/aws/data/ec2/aws_customer_gateway.json
var dataAwsCustomerGateway []byte

//go:embed mapping/aws/data/dms/aws_dms_endpoint.json
var dataAwsDmsEndpoint []byte

//go:embed mapping/aws/data/dms/aws_dms_replication_instance.json
var dataAwsDmsReplicationInstance []byte

//go:embed mapping/aws/data/dms/aws_dms_replication_subnet_group.json
var dataAwsDmsReplicationSubnetGroup []byte

//go:embed mapping/aws/data/dms/aws_dms_replication_task.json
var dataAwsDmsReplicationTask []byte

//go:embed mapping/aws/data/waf/aws_waf_ipset.json
var dataAwsWafIpset []byte

//go:embed mapping/aws/data/waf/aws_waf_rate_based_rule.json
var dataAwsWafRateBasedRule []byte

//go:embed mapping/aws/data/waf/aws_waf_rule.json
var dataAwsWafRule []byte

//go:embed mapping/aws/data/waf/aws_waf_web_acl.json
var dataAwsWafWebACL []byte

//go:embed mapping/aws/data/wafregional/aws_wafregional_ipset.json
var dataAwsWafregionalIpset []byte

//go:embed mapping/aws/data/wafregional/aws_wafregional_rate_based_rule.json
var dataAwsWafregionalRateBasedRule []byte

//go:embed mapping/aws/data/wafregional/aws_wafregional_rule.json
var dataAwsWafregionalRule []byte

//go:embed mapping/aws/data/wafregional/aws_wafregional_web_acl.json
var dataAwsWafregionalWebACL []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transitgateway_route_table_propagations.json
var dataAwsEc2TransitGatewayRouteTablePropagations []byte

//go:embed mapping/aws/data/ecr/aws_ecr_pull_through_cache_rule.json
var dataAwsEcrPullThroughCacheRule []byte

//go:embed mapping/aws/data/ecr-public/aws_ecrpublic_authorization_token.json
var dataEcrpublicAuthorizationToken []byte

//go:embed mapping/aws/data/ecs/aws_ecs_task_execution.json
var dataAwsEcsTaskExecution []byte

//go:embed mapping/aws/data/ec2/aws_eip.json
var dataAwsEip []byte

//go:embed mapping/aws/data/ec2/aws_eips.json
var dataAwsEips []byte

//go:embed mapping/aws/data/eks/aws_eks_addon.json
var dataAwsEksAddon []byte

//go:embed mapping/aws/data/eks/aws_eks_addon_version.json
var dataAwsEksAddonVersion []byte

//go:embed mapping/aws/data/fsx/aws_fsx_openzfs_snapshot.json
var dataAwsFsxOpenzfsSnapshot []byte

//go:embed mapping/aws/data/fsx/aws_fsx_windows_file_system.json
var dataAwsFsxWindowsFileSystem []byte

//go:embed mapping/aws/data/ec2/aws_instance.json
var dataAwsInstance []byte

//go:embed mapping/aws/data/ec2/aws_instances.json
var dataAwsInstances []byte

//go:embed mapping/aws/data/ec2/aws_internet_gateway.json
var dataAwsInternetGateway []byte

//go:embed mapping/aws/data/ec2/aws_key_pair.json
var dataAwsKeyPair []byte

//go:embed mapping/aws/data/kms/aws_kms_secrets.json
var dataAwsKmsSecrets []byte

//go:embed mapping/aws/data/mq/aws_mq_broker.json
var dataAwsMqBroker []byte

//go:embed mapping/aws/data/ec2/aws_nat_gateway.json
var dataAwsNatGateway []byte

//go:embed mapping/aws/data/ec2/aws_nat_gateways.json
var dataAwsNatGateways []byte

//go:embed mapping/aws/data/rds/aws_rds_clusters.json
var dataAwsRdsClusters []byte

//go:embed mapping/aws/data/ec2/aws_route.json
var dataAwsRoute []byte

//go:embed mapping/aws/data/s3/aws_s3_account_public_access_block.json
var dataAwsS3AccountPublicAccessBlock []byte

//go:embed mapping/aws/data/s3/aws_s3_bucket_policy.json
var dataAwsS3BucketPolicy []byte

//go:embed mapping/aws/data/s3/aws_s3_objects.json
var dataAwsS3Objects []byte

//go:embed mapping/aws/data/ssm/aws_ssm_document.json
var dataAwsSsmDocument []byte

//go:embed mapping/aws/data/transfer/aws_transfer_server.json
var dataAwsTransferServer []byte

//go:embed mapping/aws/data/ec2/aws_vpc_endpoint.json
var dataAwsVpcEndpoint []byte

//go:embed mapping/aws/data/sqs/aws_sqs_queue.json
var dataAwsSqsQueue []byte

//go:embed mapping/aws/data/ec2/aws_ebs_default_kms_key.json
var dataAwsEbsEncryptionByDefault []byte

//go:embed mapping/aws/data/ec2/aws_ec2_spot_price.json
var dataAwsEc2SpotPrice []byte

//go:embed mapping/aws/data/iam/aws_iam_access_keys.json
var dataAwsIamAccessKeys []byte

//go:embed mapping/aws/data/iam/aws_iam_principal_policy_simulation.json
var dataAwsIamPrincipalPolicySimulation []byte

//go:embed mapping/aws/data/kinesis/aws_kinesis_firehose_delivery_stream.json
var dataAwsKinesisFirehoseDeliveryStream []byte

//go:embed mapping/aws/data/kinesis/aws_kinesis_stream.json
var dataAwsKinesisStream []byte

//go:embed mapping/aws/data/ec2/aws_launch_template.json
var dataAwsLaunchTemplate []byte

//go:embed mapping/aws/data/ec2/aws_regions.json
var dataAwsRegions []byte

//go:embed mapping/aws/data/states/aws_sfn_state_machine.json
var dataAwsSfnStateMachine []byte

//go:embed mapping/aws/data/states/aws_sfn_state_machine_versions.json
var dataAwsSfnStateMachineVersion []byte

//go:embed mapping/aws/data/ec2/aws_vpn_gateway.json
var dataAwsVpnGateway []byte

//go:embed mapping/aws/data/lambda/aws_lambda_code_signing_config.json
var dataAwsLambdaCodeSigningConfig []byte

//go:embed mapping/aws/data/lambda/aws_lambda_invocation.json
var dataAwsLambdaInvocation []byte

//go:embed mapping/aws/data/rds/aws_neptune_engine_version.json
var dataAwsNeptuneEngineVersion []byte

//go:embed mapping/aws/data/rds/aws_neptune_orderable_db_instance.json
var dataAwsNeptuneOrderableDBInstance []byte

//go:embed mapping/aws/data/ec2/aws_network_interfaces.json
var dataAwsNetworkInterfaces []byte

//go:embed mapping/aws/data/quicksight/aws_quicksight_group.json
var dataAwsQuicksightGroup []byte

//go:embed mapping/aws/data/quicksight/aws_quicksight_theme.json
var dataAwsQuicksightTheme []byte

//go:embed mapping/aws/data/quicksight/aws_quicksight_user.json
var dataAwsQuicksightUser []byte

//go:embed mapping/aws/data/states/aws_sfn_activity.json
var dataAwsSfnActivity []byte

//go:embed mapping/aws/data/states/aws_sfn_alias.json
var dataAwsSfnAlias []byte

//go:embed mapping/aws/data/ssm/aws_ssm_instances.json
var dataAwsSsmInstances []byte

//go:embed mapping/aws/data/ssm/aws_ssm_maintenance_windows.json
var dataAwsSsmMaintenanceWindows []byte

//go:embed mapping/aws/data/ssm/aws_ssm_parameters_by_path.json
var dataAwsSsmParametersByPath []byte

//go:embed mapping/aws/data/ssm/aws_ssm_patch_baseline.json
var dataAwsSsmPatchBaseline []byte

//go:embed mapping/aws/data/ec2/aws_ec2_client_vpn_endpoint.json
var dataAwsEc2ClientVpnEndpoint []byte

//go:embed mapping/aws/data/ec2/aws_ec2_coip_pool.json
var dataAwsEc2CoipPool []byte

//go:embed mapping/aws/data/ec2/aws_ec2_coip_pools.json
var dataAwsEc2CoipPools []byte

//go:embed mapping/aws/data/ec2/aws_ec2_host.json
var dataAwsEc2Host []byte

//go:embed mapping/aws/data/ec2/aws_ec2_instance_type.json
var dataAwsEc2InstanceType []byte

//go:embed mapping/aws/data/ec2/aws_ec2_instance_type_offering.json
var dataAwsEc2InstanceTypeOffering []byte

//go:embed mapping/aws/data/ec2/aws_ec2_instance_type_offerings.json
var dataAwsEc2InstanceTypeOfferings []byte

//go:embed mapping/aws/data/ec2/aws_ec2_instance_types.json
var dataAwsEc2InstanceTypes []byte

//go:embed mapping/aws/data/ec2/aws_ec2_local_gateway.json
var dataAwsEc2LocalGateway []byte

//go:embed mapping/aws/data/ec2/aws_ec2_local_gateway_route_table.json
var dataAwsEc2LocalGatewayRouteTable []byte

//go:embed mapping/aws/data/ec2/aws_ec2_local_gateway_route_tables.json
var dataAwsEc2LocalGatewayRouteTables []byte

//go:embed mapping/aws/data/ec2/aws_ec2_local_gateway_virtual_interface.json
var dataAwsEc2LocalGatewayVirtualInterface []byte

//go:embed mapping/aws/data/ec2/aws_ec2_local_gateway_virtual_interface_group.json
var dataAwsEc2LocalGatewayVirtualInterfaceGroup []byte

//go:embed mapping/aws/data/ec2/aws_ec2_local_gateway_virtual_interface_groups.json
var dataAwsEc2LocalGatewayVirtualInterfaceGroups []byte

//go:embed mapping/aws/data/ec2/aws_ec2_local_gateways.json
var dataAwsEc2LocalGateways []byte

//go:embed mapping/aws/data/ec2/aws_ec2_managed_prefix_lists.json
var dataAwsEc2ManagedPrefixLists []byte

//go:embed mapping/aws/data/ec2/aws_ec2_public_ipv4_pool.json
var dataAwsEc2PublicIpv4Pool []byte

//go:embed mapping/aws/data/ec2/aws_ec2_public_ipv4_pools.json
var dataAwsEc2PublicIpv4Pools []byte

//go:embed mapping/aws/data/ec2/aws_ec2_serial_console_access.json
var dataAwsEc2SerialConsoleAccess []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_connect.json
var dataAwsEc2TransitGatewayConnect []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_connect_peer.json
var dataAwsEc2TransitGatewayConnectPeer []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_multicast_domain.json
var dataAwsEc2TransitGatewayMulticastDomain []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_peering_attachment.json
var dataAwsEc2TransitGatewayPeeringAttachment []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_route_table.json
var dataAwsEc2TransitGatewayRouteTable []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_route_table_associations.json
var dataAwsEc2TransitGatewayRouteTableAssociations []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_route_table_routes.json
var dataAwsEc2TransitGatewayRouteTableRoutes []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_route_tables.json
var dataAwsEc2TransitGatewayRouteTables []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_vpc_attachment.json
var dataAwsEc2TransitGatewayVpcAttachment []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_vpc_attachments.json
var dataAwsEc2TransitGatewayVpcAttachments []byte

//go:embed mapping/aws/data/ec2/aws_ec2_transit_gateway_vpn_attachment.json
var dataAwsEc2TransitGatewayVPNAttachment []byte

//go:embed mapping/aws/data/globalaccelerator/aws_globalaccelerator_accelerator.json
var dataAwsGlobalAccelerator []byte

//go:embed mapping/aws/data/globalaccelerator/aws_globalaccelerator_custom_routing_accelerator.json
var dataAwsGlobalAcceleratorCustomRoutingAccelerator []byte

//go:embed mapping/aws/data/grafana/aws_grafana_workspace.json
var dataAwsGrafanaWorkspace []byte

//go:embed mapping/aws/data/guardduty/aws_guardduty_detector.json
var dataAwsGuarddutyDetector []byte

//go:embed mapping/aws/data/guardduty/aws_guardduty_finding_ids.json
var dataAwsGuarddutyFindingIds []byte

//go:embed mapping/aws/data/identitystore/aws_identitystore_group.json
var dataAwsIdentitystoreGroup []byte

//go:embed mapping/aws/data/identitystore/aws_identitystore_user.json
var dataAwsIdentitystoreUser []byte

//go:embed mapping/aws/data/iot/aws_iot_endpoint.json
var dataAwsIotEndpoint []byte

//go:embed mapping/aws/data/kinesis/aws_kinesis_stream_consumer.json
var dataAwsKinesisStreamConsumer []byte

//go:embed mapping/aws/data/workspaces/aws_workspaces_image.json
var dataAwsWorkspaceImage []byte

//go:embed mapping/aws/data/workspaces/aws_workspaces_directory.json
var dataAwsWorkspaceDirectory []byte

//go:embed mapping/aws/data/lakeformation/aws_lakeformation_data_lake_settings.json
var dataAwsLakeformationDataLakeSettings []byte

//go:embed mapping/aws/data/lakeformation/aws_lakeformation_permissions.json
var dataAwsLakeformationPermissions []byte

//go:embed mapping/aws/data/lakeformation/aws_lakeformation_resource.json
var dataAwsLakeformationResource []byte

//go:embed mapping/aws/data/mq/aws_mq_broker_instance_type_offerings.json
var dataAwsMqBrokerInstanceTypeOfferings []byte

//go:embed mapping/aws/data/secretsmanager/aws_secretsmanager_random_password.json
var dataAwsSecretsmanagerRandomPassword []byte

//go:embed mapping/aws/data/secretsmanager/aws_secretsmanager_secret_rotation.json
var dataAwsSecretsmanagerSecretRotation []byte

//go:embed mapping/aws/data/secretsmanager/aws_secretsmanager_secrets.json
var dataAwsSecretsmanagerSecrets []byte

//go:embed mapping/aws/data/sso/aws_ssoadmin_permission_set.json
var dataAwsSsoadminPermissionSet []byte

//go:embed mapping/aws/data/ec2/aws_vpc_peering_connection.json
var dataAwsVpcPeeringConnection []byte

//go:embed mapping/aws/data/ec2/aws_vpc_peering_connections.json
var dataAwsVpcPeeringConnections []byte

//go:embed mapping/aws/data/ec2/aws_vpc_security_group_rule.json
var dataAwsVpcSecurityGroupRule []byte

//go:embed mapping/aws/data/ec2/aws_vpc_security_group_rules.json
var dataAwsVpcSecurityGroupRules []byte

//go:embed mapping/aws/data/kafka/aws_msk_vpc_connection.json
var dataAwsMskVpcConnection []byte

//go:embed mapping/aws/data/kafka/aws_mskconnect_connector.json
var dataAwsMskconnectConnector []byte

//go:embed mapping/aws/data/kafka/aws_mskconnect_custom_plugin.json
var dataAwsMskconnectCustomPlugin []byte

//go:embed mapping/aws/data/kafka/aws_mskconnect_worker_configuration.json
var dataAwsMskconnectWorkerConfiguration []byte

//go:embed mapping/aws/data/ram/aws_ram_resource_share.json
var dataAwsRAMResourceShare []byte

//go:embed mapping/aws/data/tag/aws_resourcegroupstaggingapi_resources.json
var dataAwsResourcegroupstaggingapiResources []byte

//go:embed mapping/aws/data/serverlessrepo/aws_serverlessapplicationrepository_application.json
var dataAwsServerlessapplicationrepositoryApplication []byte

//go:embed mapping/aws/data/signer/aws_signer_signing_job.json
var dataAwsSignerSigningJob []byte

//go:embed mapping/aws/data/signer/aws_signer_signing_profile.json
var dataAwsSignerSigningProfile []byte

//go:embed mapping/aws/data/ssm-incidents/aws_ssmincidents_replication_set.json
var dataAwsSsmincidentsReplicationSet []byte

//go:embed mapping/aws/data/license-manager/aws_licensemanager_grants.json
var dataAwsLicensemanagerGrants []byte

//go:embed mapping/aws/data/license-manager/aws_licensemanager_received_license.json
var dataAwsLicensemanagerReceivedLicense []byte

//go:embed mapping/aws/data/license-manager/aws_licensemanager_received_licenses.json
var dataAwsLicensemanagerReceivedLicenses []byte

//go:embed mapping/aws/data/network-firewall/aws_networkfirewall_firewall.json
var dataAwsNetworkfirewallFirewall []byte

//go:embed mapping/aws/data/network-firewall/aws_networkfirewall_firewall_policy.json
var dataAwsNetworkfirewallFirewallPolicy []byte

//go:embed mapping/aws/data/network-firewall/aws_networkfirewall_resource_policy.json
var dataAwsNetworkfirewallResourcePolicy []byte

//go:embed mapping/aws/data/qldb/aws_qldb_ledger.json
var dataAwsQldbLedger []byte

//go:embed mapping/aws/data/redshift-serverless/awas_redshiftserverless_namespace.json
var dataAwsRedshiftserverlessNamespace []byte

//go:embed mapping/aws/data/redshift-serverless/awas_redshiftserverless_workgroups.json
var dataAwsRedshiftserverlessWorkgroup []byte

//go:embed mapping/aws/data/geo/aws_location_geofence_collection.json
var dataAwsLocationGeofenceCollection []byte

//go:embed mapping/aws/data/geo/aws_location_map.json
var dataAwsLocationMap []byte

//go:embed mapping/aws/data/geo/aws_location_place_index.json
var dataAwsLocationPlaceIndex []byte

//go:embed mapping/aws/data/geo/aws_location_route_calculator.json
var dataAwsLocationRouteCalculator []byte

//go:embed mapping/aws/data/geo/aws_location_tracker.json
var dataAwsLocationTracker []byte

//go:embed mapping/aws/data/vpc-lattice/aws_vpclattice_auth_policy.json
var dataAwsVpclatticeAuthPolicy []byte

//go:embed mapping/aws/data/vpc-lattice/aws_vpclattice_listener.json
var dataAwsVpclatticeListener []byte

//go:embed mapping/aws/data/vpc-lattice/aws_vpclattice_resource_policy.json
var dataAwsVpclatticeResourcePolicy []byte

//go:embed mapping/aws/data/vpc-lattice/aws_vpclattice_service.json
var dataAwsVpclatticeService []byte

//go:embed mapping/aws/data/vpc-lattice/aws_vpclattice_service_network.json
var dataAwsVpclatticeServiceNetwork []byte

//go:embed mapping/aws/data/dms/aws_dms_certificate.json
var dataAwsDmsCertificate []byte

//go:embed mapping/aws/data/ses/aws_ses_active_receipt_rule_set.json
var dataAwsSesActiveReceiptRuleSet []byte

//go:embed mapping/aws/data/ses/aws_ses_domain_identity.json
var dataAwsSesDomainIdentity []byte

//go:embed mapping/aws/data/ses/aws_ses_email_identity.json
var dataAwsSesEmailIdentity []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_component.json
var dataAwsImagebuilderComponent []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_components.json
var dataAwsImagebuilderComponets []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_container_recipe.json
var dataAwsImagebuilderContainerRecipe []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_container_recipes.json
var dataAwsImagebuilderContainerRecipes []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_distribution_configuration.json
var dataAwsImagebuilderDistributionConfiguration []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_distribution_configurations.json
var dataAwsImagebuilderDistributionConfigurations []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_image.json
var dataAwsImagebuilderImage []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_image_pipeline.json
var dataAwsImagebuilderImagePipeline []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_image_pipelines.json
var dataAwsImagebuilderImagePipelines []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_image_recipe.json
var dataAwsImagebuilderImageRecipe []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_image_recipes.json
var dataAwsImagebuilderImageRecipes []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_infrastructure_configuration.json
var dataAwsImagebuilderInfrastructureConfiguration []byte

//go:embed mapping/aws/data/imagebuilder/aws_imagebuilder_distribution_configurations.json
var dataAwsImagebuilderInfrastructureConfigurations []byte

//go:embed mapping/aws/data/kendra/aws_kendra_experience.json
var dataAwsKendraExperience []byte

//go:embed mapping/aws/data/kendra/aws_kendra_faq.json
var dataAwsKendraFaq []byte

//go:embed mapping/aws/data/kendra/aws_kendra_index.json
var dataAwsKendraIndex []byte

//go:embed mapping/aws/data/kendra/aws_kendra_query_suggestions_block_list.json
var dataAwsKendraQuerySuggestionsBlockList []byte

//go:embed mapping/aws/data/kendra/aws_kendra_thesaurus.json
var dataAwsKendraThesaurus []byte

//go:embed mapping/aws/data/lex/aws_lex_bot.json
var dataAwsLexBot []byte

//go:embed mapping/aws/data/lex/aws_lex_bot_alias.json
var dataAwsLexBotAlias []byte

//go:embed mapping/aws/data/lex/aws_lex_intent.json
var dataAwsLexIntent []byte

//go:embed mapping/aws/data/lex/aws_lex_slot_type.json
var dataAwsLexSlotType []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_connection.json
var dataAwsNetworkManagerConnection []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_connections.json
var dataAwsNetworkManagerConnections []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_device.json
var dataAwsNetworkManagerDevice []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_devices.json
var dataAwsNetworkManagerDevices []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_global_network.json
var dataAwsNetworkManagerGlobalNetwork []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_global_networks.json
var dataAwsNetworkManagerGlobalNetworks []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_link.json
var dataAwsNetworkManagerLink []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_links.json
var dataAwsNetworkManagerLinks []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_site.json
var dataAwsNetworkManagerSite []byte

//go:embed mapping/aws/data/networkmanager/aws_networkmanager_sites.json
var dataAwsNetworkManagerSites []byte

//go:embed mapping/aws/data/oam/aws_oam_link.json
var dataAwsOamLink []byte

//go:embed mapping/aws/data/oam/aws_oam_links.json
var dataAwsOamLinks []byte

//go:embed mapping/aws/data/oam/aws_oam_sink.json
var dataAwsOamSink []byte

//go:embed mapping/aws/data/oam/aws_oam_sinks.json
var dataAwsOamSinks []byte

//go:embed mapping/aws/data/organizations/aws_organizations_delegated_administrators.json
var dataAwsOrganizationsDelegatedAdministrators []byte

//go:embed mapping/aws/data/organizations/aws_organizations_delegated_services.json
var dataAwsOrganizationsDelegatedServices []byte

//go:embed mapping/aws/data/organizations/aws_organizations_organizational_unit_child_accounts.json
var dataAwsOrganizationsOrganizationalUnitChildsAccounts []byte

//go:embed mapping/aws/data/organizations/aws_organizations_organizational_unit_descendant_accounts.json
var dataAwsOrganizationsOrganizationalUnitDescendantAccounts []byte

//go:embed mapping/aws/data/organizations/aws_organizations_organizational_units.json
var dataAwsOrganizationsOrganizationalUnits []byte

//go:embed mapping/aws/data/organizations/aws_organizations_policies.json
var dataAwsOrganizationsPolicies []byte

//go:embed mapping/aws/data/organizations/aws_organizations_policies_for_target.json
var dataAwsOrganizationsPoliciesForTarget []byte

//go:embed mapping/aws/data/organizations/aws_organizations_resource_tags.json
var dataAwsOrganizationsResourceTags []byte

//go:embed mapping/aws/data/outposts/aws_outposts_asset.json
var dataAwsOutpostsAsset []byte

//go:embed mapping/aws/data/outposts/aws_outposts_assets.json
var dataAwsOutpostsAssets []byte

//go:embed mapping/aws/data/outposts/aws_outposts_outpost_instance_type.json
var dataAwsOutpostsOutpostInstanceType []byte

//go:embed mapping/aws/data/outposts/aws_outposts_outpost_instance_types.json
var dataAwsOutpostsOutpostInstanceTypes []byte

//go:embed mapping/aws/data/outposts/aws_outposts_sites.json
var dataAwsOutpostSites []byte

//go:embed mapping/aws/data/outposts/aws_outposts_sites.json
var dataAwsOutpostsSites []byte

//go:embed mapping/aws/data/route53/aws_route53_delegation_set.json
var dataAwsRoute53DelegationSet []byte

//go:embed mapping/aws/data/route53resolver/aws_route53_resolver_endpoint.json
var dataAwsRoute53ResolverEndpoint []byte

//go:embed mapping/aws/data/route53resolver/aws_route53_resolver_query_log_config.json
var dataAwsRoute53ResolverQueryLogConfig []byte

//go:embed mapping/aws/data/route53resolver/aws_route53_resolver_rules.json
var dataAwsRoute53ResolverRules []byte

//go:embed mapping/aws/data/servicecatalog/aws_servicecatalog_constraint.json
var dataAwsSevicecatalogConstraint []byte

//go:embed mapping/aws/data/servicecatalog/aws_servicecatalog_launch_paths.json
var dataAwsSevicecatalogLaunchPaths []byte

//go:embed mapping/aws/data/servicecatalog/aws_servicecatalog_portfolio.json
var dataAwsSevicecatalogPortfolio []byte

//go:embed mapping/aws/data/servicecatalog/aws_servicecatalog_portfolio_constraints.json
var dataAwsSevicecatalogPortfolioConstraints []byte

//go:embed mapping/aws/data/servicecatalog/aws_servicecatalog_product.json
var dataAwsSevicecatalogProduct []byte

//go:embed mapping/aws/data/servicecatalog/aws_servicecatalog_provisioning_artifacts.json
var dataAwsSevicecatalogProvisioningArtifacts []byte

//go:embed mapping/aws/data/ses/aws_sesv2_configuration_set.json
var dataAwsSesv2ConfigurationSet []byte

//go:embed mapping/aws/data/ses/aws_ses_email_identity.json
var dataAwsSesv2EmailIdentity []byte

//go:embed mapping/aws/data/aoss/aws_opensearchserverless_access_policy.json
var dataAwsOpensearchserverlessAccessPolicy []byte

//go:embed mapping/aws/data/aoss/aws_opensearchserverless_collection.json
var dataAwsOpensearchserverlessCollection []byte

//go:embed mapping/aws/data/aoss/aws_opensearchserverless_security_config.json
var dataAwsOpensearchserverlessSecurityConfig []byte

//go:embed mapping/aws/data/aoss/aws_opensearchserverless_security_policy.json
var dataAwsOpensearchserverlessSecurityPolicy []byte

//go:embed mapping/aws/data/aoss/aws_opensearchserverless_vpc_endpoint.json
var dataAwsOpensearchserverlessVpcEndpoint []byte

//go:embed mapping/aws/data/cognito-identity/aws_cognito_identity_pool.json
var dataAwsCognitoIdentityPool []byte

//go:embed mapping/aws/data/fsx/aws_fsx_ontap_file_system.json
var dataAwsFsxOntapFileSystem []byte

//go:embed mapping/aws/data/fsx/aws_fsx_ontap_storage_virtual_machine.json
var dataAwsFsxOntapStorageVirtualMachine []byte

//go:embed mapping/aws/data/fsx/aws_fsx_ontap_storage_virtual_machines.json
var dataAwsFsxOntapStorageVirtualMachines []byte

//go:embed mapping/aws/data/organizations/aws_organizations_organizational_unit.json
var dataAwsOrganizationsOrganizationalUnit []byte

//go:embed mapping/aws/data/organizations/aws_organizations_policy.json
var dataAwsOrganizationsPolicy []byte

//go:embed mapping/aws/data/servicequota/aws_servicequotas_templates.json
var dataAwsServicequotasTemplates []byte

//go:embed mapping/aws/data/apigateway/aws_apigatewayv2_vpc_link.json
var dataAwsApigatewayv2VpcLink []byte

//go:embed mapping/aws/data/athena/aws_athena_named_query.json
var dataAwsAthenaNamedQuery []byte

//go:embed mapping/aws/data/bedrock/aws_bedrock_foundation_model.json
var dataAwsBedrockFoundationModel []byte

//go:embed mapping/aws/data/bedrock/aws_bedrock_foundation_models.json
var dataAwsBedrockFoundationModels []byte

//go:embed mapping/aws/data/iot/aws_iot_registration_code.json
var dataAwsIotRegistrationCode []byte

//go:embed mapping/aws/data/aoss/aws_opensearchserverless_lifecycle_policy.json
var dataAwsOpensearchserverlessLifecyclePolicy []byte

//go:embed mapping/aws/data/elasticmapreduce/aws_emr_supported_instance_types.json
var dataAwsEmrSupportedInstanceTypes []byte

//go:embed mapping/aws/data/elasticloadbalancing/aws_lb_trust_store.json
var dataAwsLbTrustStore []byte

//go:embed mapping/aws/data/codeguru-profiler/aws_codeguruprofiler_profiling_group.json
var dataAwsCodeguruprofilerProfilingGroup []byte

//go:embed mapping/aws/data/ecr/aws_ecr_repositories.json
var dataAwsEcrRepositories []byte

//go:embed mapping/aws/data/sso/aws_ssoadmin_application_providers.json
var dataAwsSsoadminApplicationProviders []byte

//go:embed mapping/aws/data/polly/aws_polly_voices.json
var dataAwsPollyVoices []byte

//go:embed mapping/aws/data/s3express/aws_s3_directory_buckets.json
var dataAwsS3DirectoryBuckets []byte

//go:embed mapping/aws/data/sso/aws_ssoadmin_application_assignments.json
var dataAwsSsoadminApplicationAssignments []byte

//go:embed mapping/aws/data/sso/aws_ssoadmin_principal_application_assignments.json
var dataAwsSsoadminPrincipalApplicationAssignments []byte

//go:embed mapping/aws/data/verifiedpermissions/aws_verifiedpermissions_policy_store.json
var dataVerifiedpermissionsPolicyStore []byte

//go:embed mapping/aws/data/kafka/aws_msk_bootstrap_brokers.json
var dataAwsMskBoostrapBrokers []byte

//go:embed mapping/aws/data/mq/aws_mq_broker_engine_types.json
var dataMqBrokerEngineTypes []byte

//go:embed mapping/aws/data/eks/aws_eks_access_entry.json
var dataAwsEksAccessEntry []byte

//go:embed mapping/aws/data/bedrock/aws_bedrock_custom_model.json
var dataAwsBedrockCustomModel []byte

//go:embed mapping/aws/data/bedrock/aws_bedrock_custom_models.json
var dataAwsBedrockCustomModels []byte

//go:embed mapping/aws/data/ssm-contacts/aws_ssmcontacts_rotation.json
var dataAwsSsmcontactsRotation []byte

//go:embed mapping/aws/data/batch/aws_batch_job_definition.json
var dataAwsBatchJobDefinition []byte

//go:embed mapping/aws/data/cognito-idp/aws_cognito_user_group.json
var dataAwsCognitoUserGroup []byte

//go:embed mapping/aws/data/cognito-idp/aws_cognito_user_groups.json
var dataAwsCognitoUserGroups []byte

//go:embed mapping/aws/data/rds/aws_db_parameter_group.json
var dataAwsDbParameterGroup []byte

//go:embed mapping/aws/data/medialive/aws_medialive_input.json
var dataAwsMedialiveInput []byte

//go:embed mapping/aws/data/redshift/aws_redshift_data_shares.json
var dataAwsRedshiftDataShares []byte

//go:embed mapping/aws/data/redshift/aws_redshift_producer_data_shares.json
var dataAwsRedshiftProducerDataShares []byte

//go:embed mapping/aws/data/resource-explorer-2/aws_resourceexplorer2_search.json
var dataAwsResourceexplorer2Search []byte

//go:embed mapping/aws/data/servicecatalog/aws_servicecatalogappregistry_application.json
var dataAwsServicecatalogappregistryApplication []byte

//go:embed mapping/aws/data/devopsguru/aws_devopsguru_notification_channel.json
var dataAwsDevopsguruNotificationChannel []byte

//go:embed mapping/aws/data/devopsguru/aws_devopsguru_resource_collection.json
var dataAwsDevopsguruResourceCollection []byte

//go:embed mapping/aws/data/sso/aws_identitystore_groups.json
var dataAwsIdentitystoreGroups []byte

//go:embed mapping/aws/data/datazone/aws_datazone_environment_blueprint.json
var dataAwsDatazoneEnvironmentBlueprint []byte

//go:embed mapping/aws/data/chatbot/aws_chatbot_slack_workspace.json
var dataAwsChatbotSlackWorkspace []byte

//go:embed mapping/aws/data/ec2/aws_ec2_capacity_block_offering.json
var dataAwsEc2CapacityBlockOffering []byte
