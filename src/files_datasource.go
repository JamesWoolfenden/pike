package pike

import (
	_ "embed" // required for embed
)

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

//go:embed mapping/aws/data/ec2/aws_vpc_endpoint_service.json
var dataAwsVpcEndpointService []byte
