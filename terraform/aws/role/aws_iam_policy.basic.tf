resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:DescribeNetworkInterfaces",
          "ec2:CreateSecurityGroup",
          "ec2:DetachNetworkInterface",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeRegions",
          "ec2:DescribeSecurityGroups",
          "cloudfront:CreateDistribution",
          "cloudfront:DeleteDistribution",
          "cloudfront:GetDistribution",
          "cloudfront:ListTagsForResource",
          "cloudfront:TagResource",
          "cloudfront:UntagResource",
          "cloudfront:UpdateDistribution",
          "ecr:CreateRepository",
          "ecr:DeleteRepository",
          "ecr:DeleteRepositoryPolicy",
          "ecr:DescribeRepositories",
          "ecr:GetRepositoryPolicy",
          "ecr:ListTagsForResource",
          "ecr:SetRepositoryPolicy",
          "elasticloadbalancing:EnableAvailabilityZonesForLoadBalancer",
          "elasticloadbalancing:AttachLoadBalancerToSubnets",
          "elasticloadbalancing:CreateLoadBalancer",
          "elasticloadbalancing:CreateLoadBalancerListeners",
          "elasticloadbalancing:DeleteLoadBalancer",
          "elasticloadbalancing:DescribeLoadBalancerAttributes",
          "elasticloadbalancing:DescribeLoadBalancers",
          "elasticloadbalancing:DescribeTags",
          "elasticloadbalancing:ModifyLoadBalancerAttributes",
          "elasticmapreduce:DescribeCluster",
          "elasticmapreduce:GetAutoTerminationPolicy",
          "elasticmapreduce:ListBootstrapActions",
          "elasticmapreduce:ListSteps",
          "elasticmapreduce:RunJobFlow",
          "elasticmapreduce:TerminateJobFlows",
          "iam:CreateServiceLinkedRole",
          "iam:PassRole",
          "s3:ListBucket",
          "s3:GetObject",
          "s3:GetBucket",

          //aws_ecr_replication_configuration
          "ecr:PutReplicationConfiguration",
          "ecr:DescribeRegistry",

          //aws_ecr_registry_policy
          "ecr:PutRegistryPolicy",
          "ecr:GetRegistryPolicy",
          "ecr:DeleteRegistryPolicy",

          //aws_codebuild_report_group
          "codebuild:CreateReportGroup",
          "codebuild:BatchGetReportGroups",
          "codebuild:DeleteReportGroup",

          //aws_codebuild_resource_policy
          "codebuild:PutResourcePolicy",
          "codebuild:GetResourcePolicy",
          "codebuild:DeleteResourcePolicy",

          //aws_cloudfront_origin_request_policy
          "cloudfront:CreateOriginRequestPolicy",
          "cloudfront:GetOriginRequestPolicy",
          "cloudfront:DeleteOriginRequestPolicy",

          //aws_cloudfront_cache_policy
          "cloudfront:CreateCachePolicy",
          "cloudfront:GetCachePolicy",
          "cloudfront:DeleteCachePolicy",

          //aws_lb_cookie_stickiness_policy
          "elasticloadbalancing:CreateLBCookieStickinessPolicy",
          "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",

          //aws_cloudfront_continuous_deployment_policy
          "cloudfront:CreateContinuousDeploymentPolicy",
          "cloudfront:GetContinuousDeploymentPolicy",
          "cloudfront:DeleteContinuousDeploymentPolicy",

          //aws_app_cookie_stickiness_policy
          "elasticloadbalancing:CreateAppCookieStickinessPolicy",
          "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
          "elasticloadbalancing:DescribeLoadBalancerPolicies",
          "elasticloadbalancing:DeleteLoadBalancerPolicy",
        ],
        "Resource" : "*",
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
