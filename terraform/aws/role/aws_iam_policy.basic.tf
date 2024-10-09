resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          //aws_cloudfrontkeyvaluestore_key
          "cloudfront-keyvaluestore:DescribeKeyValueStore",
          "cloudfront-keyvaluestore:PutKey",
          "cloudfront-keyvaluestore:GetKey",

          //aws-ecs_tag
          "ecs:TagResource",
          "ecs:UntagResource",

          //aws_lb_trust_store
          "elasticloadbalancing:CreateTrustStore",
          "elasticloadbalancing:DeleteTrustStore",
          "elasticloadbalancing:ModifyTrustStore",

          //aws_quicksight_folder
          "quicksight:CreateFolder",
          "quicksight:DescribeFolder",
          "quicksight:DeleteFolder",
          "quicksight:UpdateFolder",

          //aws_quicksight_user
          "quicksight:RegisterUser",
          "quicksight:DescribeUser",
          "quicksight:CreateUser",
          "quicksight:DeleteUser",
          "quicksight:UpdateUser",

          //aws_quicksight_namespace
          "quicksight:CreateNamespace",
          "quicksight:DeleteNamespace",

          //aws_quicksight_group
          "quicksight:CreateGroup",
          "quicksight:DescribeGroup",
          "quicksight:DeleteGroup",
          "quicksight:UpdateGroup",

          //aws_quicksight_group_membership
          "quicksight:DescribeGroupMembership",
          "quicksight:CreateGroupMembership",
          "quicksight:DeleteGroupMembership",

          "batch:CreateComputeEnvironment",
          "batch:DeleteComputeEnvironment",
          "batch:DescribeComputeEnvironments",
          "batch:TagResource",
          "batch:UntagResource",
          "batch:UpdateComputeEnvironment"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor1",
        "Effect" : "Allow",
        "Action" : [
          "cloudfront:CreateKeyValueStore",
          "cloudfront:DeleteKeyValueStore",
          "cloudfront:DescribeKeyValueStore",
          "cloudfront:UpdateKeyValueStore"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor2",
        "Effect" : "Allow",
        "Action" : [
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor3",
        "Effect" : "Allow",
        "Action" : [
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeImages",
          "ec2:DescribeKeyPairs",
          "ec2:DescribeLaunchTemplateVersions",
          "ec2:DescribeLaunchTemplates",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeVpcs"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor4",
        "Effect" : "Allow",
        "Action" : [
          "ecs:Describe*",
          "ecs:DescribeClusters",
          "ecs:List*"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor5",
        "Effect" : "Allow",
        "Action" : [
          "iam:CreateServiceLinkedRole",
          "iam:PassRole"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor6",
        "Effect" : "Allow",
        "Action" : [
          "s3:CreateBucket",
          "s3:DeleteBucket",
          "s3:DeleteObject",
          "s3:GetAccelerateConfiguration",
          "s3:GetBucketAcl",
          "s3:GetBucketCORS",
          "s3:GetBucketLogging",
          "s3:GetBucketObjectLockConfiguration",
          "s3:GetBucketPolicy",
          "s3:GetBucketRequestPayment",
          "s3:GetBucketTagging",
          "s3:GetBucketVersioning",
          "s3:GetBucketWebsite",
          "s3:GetEncryptionConfiguration",
          "s3:GetLifecycleConfiguration",
          "s3:GetObject",
          "s3:GetObjectAcl",
          "s3:GetReplicationConfiguration",
          "s3:ListBucket",
          "s3:PutObject"
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
