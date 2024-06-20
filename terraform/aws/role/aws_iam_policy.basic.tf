resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          //aws_dms_event_subscription
          "dms:DescribeEventSubscriptions",
          "dms:CreateEventSubscription",
          "dms:DeleteEventSubscription",
          "dms:ModifyEventSubscription",

          //aws_dms_replication_config
          "dms:CreateReplicationConfig",
          "dms:DescribeReplicationConfigs",
          "dms:DeleteReplicationConfig",
          "dms:ModifyReplicationConfig",

          //aws_dms_replication_instance
          "dms:DescribeReplicationInstances",
          "dms:CreateReplicationInstance",
          "dms:DeleteReplicationInstance",
          "dms:ModifyReplicationInstance",

          //aws_dms_replication_subnet_group
          "dms:DescribeReplicationSubnetGroups",
          "dms:CreateReplicationSubnetGroup",
          "dms:DeleteReplicationSubnetGroup",
          "dms:ModifyReplicationSubnetGroup",

          //aws_dms_replication_task
          "dms:DeleteReplicationTask",
          "dms:ModifyReplicationTask",
          "dms:CreateReplicationTask",

          //aws_dms_s3_endpoint
          "dms:DescribeEndpoints",
          "dms:CreateEndpoint",
          "dms:DeleteEndpoint",
          "dms:ModifyEndpoint",
          "iam:PassRole",



          "ec2:DescribeCapacityBlockOfferings",
          "acm:DescribeCertificate",
          "acm:ListCertificates",
          "acm:GetCertificate",
          "acm:ListTagsForCertificate",
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "iam:AttachRolePolicy",
          "iam:CreateRole",
          "iam:DeleteRole",
          "iam:DeleteRolePolicy",
          "iam:DetachRolePolicy",
          "iam:GetRole",
          "iam:GetRolePolicy",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",
          "iam:PutRolePolicy",
          "kms:DescribeKey",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",
          "sns:CreateTopic",
          "sns:DeleteTopic",
          "sns:GetTopicAttributes",
          "sns:ListTagsForResource",
          "sns:SetTopicAttributes"
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
