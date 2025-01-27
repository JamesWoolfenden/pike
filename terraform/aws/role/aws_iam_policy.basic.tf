resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [

          "config:DeleteConfigRule",
          "config:DescribeComplianceByConfigRule",
          "config:DescribeConfigRules",
          "config:ListTagsForResource",
          "config:PutConfigRule",
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeImages",
          "ec2:DescribeInstanceAttribute",
          "ec2:DescribeInstanceCreditSpecifications",
          "ec2:DescribeInstanceTypes",
          "ec2:DescribeInstances",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribeTags",
          "ec2:DescribeVolumes",
          "ec2:ModifyInstanceAttribute",
          "ec2:RunInstances",
          "ec2:StartInstances",
          "ec2:StopInstances",
          "ec2:TerminateInstances",
          "iam:AttachRolePolicy",
          "iam:CreateRole",
          "iam:CreateServiceLinkedRole",
          "iam:DeleteRole",
          "iam:DetachRolePolicy",
          "iam:GetRole",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",
          "organizations:CreateOrganization",
          "organizations:DeleteOrganization",
          "organizations:DescribeOrganization",
          "organizations:ListRoots",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",

          # aws_lakeformation_resource
          "lakeformation:RegisterResource",
          "iam:PutRolePolicy",
          "lakeformation:DescribeResource",
          "lakeformation:DeregisterResource",
          "iam:GetRolePolicy",
          "lakeformation:UpdateResource",

          # aws_lakeformation_data_lake_settings
          "lakeformation:PutDataLakeSettings",
          "lakeformation:GetDataLakeSettings",

          # aws_glue_dev_endpoint
          "glue:CreateDevEndpoint",
          "iam:PassRole",
          "glue:GetDevEndpoint",
          "glue:DeleteDevEndpoint",
          "glue:UpdateDevEndpoint",

          # aws_glue_data_quality_ruleset
          "glue:CreateDataQualityRuleset",
          "glue:GetDataQualityRuleset",
          "glue:GetTags",
          "glie:DeleteTags",
          "glue:DeleteDataQualityRuleset",
          "glue:UpdateDataQualityRuleset",


          # aws_ec2_traffic_mirror_target
          "ec2:CreateTrafficMirrorTarget",
          "ec2:DescribeTrafficMirrorTargets",
          "ec2:DeleteTrafficMirrorTarget",

          # aws_ec2_traffic_mirror_filter
          "ec2:CreateTrafficMirrorFilter",
          "ec2:ModifyTrafficMirrorFilterNetworkServices",
          "ec2:DescribeTrafficMirrorFilters",
          "ec2:DeleteTrafficMirrorFilter",

          # aws_config_aggregate_authorization
          "config:PutAggregationAuthorization",
          "config:DescribeAggregationAuthorizations",
          "config:DeleteAggregationAuthorization",

          # aws_config_organization_managed_rule
          "config:DescribeOrganizationConfigRules",
          "config:DeleteOrganizationConfigRule",
          "config:PutOrganizationConfigRule",

          # aws_config_remediation_configuration
          "config:DescribeRemediationConfigurations",
          "config:DeleteRemediationConfiguration",
          "config:PutRemediationConfigurations",

          # aws_ec2_traffic_mirror_session.tf
          "ec2:CreateTrafficMirrorSession",
          "ec2:DeleteTrafficMirrorSession",
          "ec2:ModifyTrafficMirrorSession",
          "ec2:DescribeTrafficMirrorSessions",


        ],
        "Resource" : [
          "*"
        ]
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
