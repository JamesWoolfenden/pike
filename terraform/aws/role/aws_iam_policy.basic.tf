resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          "dynamodb:CreateTable",
          "dynamodb:DeleteItem",
          "dynamodb:DeleteTable",
          "dynamodb:DescribeContinuousBackups",
          "dynamodb:DescribeTable",
          "dynamodb:DescribeTimeToLive",
          "dynamodb:GetItem",
          "dynamodb:ListTagsOfResource",
          "dynamodb:PutItem",
          "dynamodb:UpdateTable",
          "dynamodb:UpdateTimeToLive",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeAddresses",
          "ec2:DescribeAvailabilityZones",
          "ec2:DescribeInternetGateways",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeVpcAttribute",
          "iam:PassRole",
          "redshift:GetResourcePolicy",
          "redshift:PutResourcePolicy",
          "redshift-serverless:CreateNamespace",
          "redshift-serverless:CreateSnapshotCopyConfiguration",
          "redshift-serverless:CreateWorkgroup",
          "redshift-serverless:DeleteNamespace",
          "redshift-serverless:DeleteWorkgroup",
          "redshift-serverless:GetNamespace",
          "redshift-serverless:GetWorkgroup",
          "redshift-serverless:ListSnapshotCopyConfigurations",
          "redshift-serverless:ListTagsForResource",
          "redshift-serverless:UpdateNamespace",
          "redshift-serverless:UpdateWorkgroup",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",
          "kms:CreateKey",
          "kms:DescribeKey",
          "kms:GetKeyPolicy",
          "kms:GetKeyRotationStatus",
          "kms:ListResourceTags",
          "kms:PutKeyPolicy",
          "kms:ScheduleKeyDeletion",
          "kms:UpdateKeyDescription",

          //aws_dynamodb_table for
          # point_in_time_recovery {
          #   enabled = true
          # }
          "dynamodb:UpdateContinuousBackups",

          # "aws_redshift_integration"
          "redshift:CreateIntegration",
          "redshift:CreateTags",
          "redshift:DeleteTags",
          "redshift:DescribeIntegrations",
          "redshift:DeleteIntegration",
          "redshift:ModifyIntegration"

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
