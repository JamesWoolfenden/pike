package pike

import (
	"fmt"
)

// GetAWSDataPermissions gets permissions required for datasource's
func GetAWSDataPermissions(result ResourceV2) ([]string, error) {

	TFLookup := map[string]interface{}{
		"aws_vpcs":                                  dataAwsVpcs,
		"aws_subnet_ids":                            dataAwsSubnetIds,
		"aws_subnet":                                dataAwsSubnetIds,
		"aws_subnets":                               dataAwsSubnetIds,
		"aws_ami":                                   dataAwsAmi,
		"aws_iam_policy":                            dataAwsIamPolicy,
		"aws_iam_role":                              dataAwsIamRole,
		"aws_s3_bucket":                             dataAwsS3Bucket,
		"aws_vpc":                                   dataAwsVpc,
		"aws_availability_zones":                    dataAwsAvailabilityZones,
		"aws_caller_identity":                       placeholder,
		"aws_iam_policy_document":                   placeholder,
		"aws_region":                                placeholder,
		"aws_canonical_user_id":                     placeholder,
		"aws_route53_traffic_policy_document":       placeholder,
		"aws_cloudtrail_service_account":            placeholder,
		"aws_partition":                             placeholder,
		"aws_inspector_rules_packages":              dataAwsInspectorRulesPackages,
		"aws_route53_zone":                          dataAwsRoute53Zone,
		"aws_kms_ciphertext":                        dataAwsKmsCiphertext,
		"aws_kms_key":                               dataAwsKmsKey,
		"aws_ebs_default_kms_key":                   dataAwsEbsDefaultKmsKey,
		"aws_security_group":                        dataAwsSecurityGroup,
		"aws_security_groups":                       dataAwsSecurityGroup,
		"aws_sns_topic":                             dataAwsSnsTopic,
		"aws_ssm_parameter":                         dataAwsSsmParameter,
		"aws_route_tables":                          dataAwsRouteTables,
		"aws_elastic_beanstalk_solution_stack":      dataAwsElasticBeanstalkSolutionStack,
		"aws_ssoadmin_instances":                    dataAwsSsoadminInstances,
		"aws_organizations_organization":            dataAwsOrganizationsOrganization,
		"aws_s3_bucket_object":                      placeholder,
		"aws_s3_object":                             placeholder,
		"aws_arn":                                   placeholder,
		"aws_default_tags":                          placeholder,
		"aws_wafv2_ip_set":                          dataAwsWafv2IpSet,
		"aws_wafv2_regex_pattern_set":               dataAwsWafv2RegexPatternSet,
		"aws_wafv2_rule_group":                      dataAwsWafv2RuleGroup,
		"aws_wafv2_web_acl":                         dataAwsWafv2WebACL,
		"aws_cloudwatch_log_group":                  dataAwsCloudwatchLogGroup,
		"aws_directory_service_directory":           dataAwsDirectoryServiceDirectory,
		"aws_ecr_authorization":                     dataAwsEcrAuthorization,
		"aws_ecs_task_definition":                   dataAwsEcsTaskDefinition,
		"aws_eks_cluster":                           dataAwsEksCluster,
		"aws_elb_service_account":                   placeholder,
		"aws_ecr_authorization_token":               dataAwsEcrAuthorizationToken,
		"aws_lambda_function":                       dataAwsLambdaFunction,
		"aws_outposts_outpost":                      dataAwsOutpostsOutpost,
		"aws_vpc_endpoint_service":                  dataAwsVpcEndpointService,
		"aws_cognito_user_pools":                    dataAwsCognitoUserPools,
		"aws_cognito_user_pool_client":              dataAwsCognitoUserPoolClient,
		"aws_cognito_user_pool_clients":             dataAwsCognitoUserPoolClients,
		"aws_cognito_user_pool_signing_certificate": dataAwsCognitoUserPoolSigningCertificate,
		"aws_redshift_cluster":                      dataAwsRedshiftCluster,
		"aws_redshift_cluster_credentials":          dataAwsRedshiftClusterCredentials,
		"aws_redshift_orderable_cluster":            dataAwsRedshiftOrderableCluster,
		"aws_redshift_service_account":              placeholder,
		"aws_redshift_subnet_group":                 dataAwsRedshiftSubnetGroup,
		"aws_db_proxy":                              placeholder,
		"aws_db_cluster_snapshot":                   dataAwsDbClusterSnapshot,
		"aws_db_event_categories":                   dataAwsDbEventCategories,
		"aws_db_instance":                           dataAwsDbInstance,
		"aws_db_snapshot":                           dataAwsDbSnapshot,
		"aws_db_subnet_group":                       dataAwsDbSubnetGroup,
		"aws_rds_certificate":                       dataAwsRdsCertificate,
		"aws_rds_cluster":                           dataAwsRdsCluster,
		"aws_rds_engine_version":                    dataAwsRdsEngineVersion,
		"aws_rds_orderable_db_instance":             dataAwsRdsOrderableDbInstance,
		"aws_glue_connection":                       dataAwsGlueConnection,
		"aws_glue_data_catalog_encryption_settings": dataAwsDataCatalogEncryptionSettings,
		"aws_glue_script":                           dataAwsGlueScript,
		"aws_efs_access_point":                      dataAwsEfsAccessPoint,
		"aws_efs_access_points":                     dataAwsEfsAccessPoints,
		"aws_efs_file_system":                       dataAwsEfsFileSystem,
		"aws_efs_mount_target":                      dataAwsEfsMountTarget,
		"aws_iam_session_context":                   placeholder,
		"aws_iam_account_alias":                     dataAwsIamAccountAlias,
		"aws_iam_group":                             dataAwsIamGroup,
		"aws_iam_instance_profile":                  dataAwsIamInstanceProfile,
		"aws_iam_instance_profiles":                 dataAwsIamInstanceProfiles,
		"aws_iam_openid_connect_provider":           dataAwsIamOpenIDConnectProvider,
		"aws_iam_saml_provider":                     dataAwsIamSamlProvider,
		"aws_iam_server_certificate":                dataAwsIamServerCertificate,
		"aws_iam_user":                              dataAwsIamUser,
		"aws_iam_user_ssh_key":                      dataAwsIamUserSSHKey,
		"aws_iam_users":                             dataAwsIamUsers,
	}

	var Permissions []string
	var err error

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}
