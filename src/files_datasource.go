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

//go:embed mapping/aws/data/logs/aws_cloudwatch_log_group.json
var dataAwsCloudwatchLogGroup []byte

//go:embed mapping/aws/data/directoryservice/aws_directory_service_directory.json
var dataAwsDirectoryServiceDirectory []byte

//go:embed mapping/aws/data/sso/aws_ssoadmin_instances.json
var dataAwsSsoadminInstances []byte

//go:embed mapping/aws/data/ecr/aws_ecr_authorization_token.json
var dataAwsEcrAuthorization []byte

//go:embed mapping/aws/data/ecs/aws_ecs_task_definition.json
var dataAwsEcsTaskDefinition []byte

//go:embed mapping/aws/data/eks/aws_eks_cluster.json
var dataAwsEksCluster []byte

//go:embed mapping/aws/data/lambda/aws_lambda_function.json
var dataAwsLambdaFunction []byte

//go:embed mapping/aws/data/ecr/aws_ecr_authorization_token.json
var dataAwsEcrAuthorizationToken []byte

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
