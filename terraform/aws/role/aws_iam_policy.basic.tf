resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          //aws_prometheus_workspace
          //placeholder

          //aws_prometheus_query_logging_configuration.example
          //placeholder

          //aws_cloudwatch_log_group.example
          //placeholder

          //aws_billing_views
          "billing:ListBillingViews",

          //aws_media_convert_queue
          "mediaconvert:GetQueue",
          "mediaconvert:ListTagsForResource",

          //aws_memorydb_acl
          "memorydb:DescribeACLs",

          //aws_memorydb_cluster
          "memorydb:DescribeClusters",

          //aws_memorydb_parameter_group
          "memorydb:DescribeParameterGroups",

          //aws_memorydb_snapshot
          "memorydb:DescribeSnapshots",

          //aws_memorydb_subnet_group
          "memorydb:DescribeSubnetGroups",
          "memorydb:ListTags",

          //aws_memorydb_user
          "memorydb:DescribeUser",
          "memorydb:DescribeUsers",

          //aws_securityhub_standards_control_associations
          "securityhub:ListStandardsControlAssociations",
          "securityhub:DescribeStandardsControls",

          //aws_workspaces_workspace
          "workspaces:DescribeWorkspaces"
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
