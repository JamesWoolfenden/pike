resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "elasticfilesystem:DescribeMountTargets",
          "elasticfilesystem:DescribeMountTargetSecurityGroups",

          "elasticfilesystem:DescribeReplicationConfigurations",
          "elasticfilesystem:DeleteReplicationConfiguration",
          "ec2:DescribeAccountAttributes",
          "elasticfilesystem:CreateReplicationConfiguration",


          "elasticfilesystem:CreateFileSystem",
          #          "elasticfilesystem:DescribeFileSystems",
          #          "elasticfilesystem:DescribeLifecycleConfiguration",
          #          "elasticfilesystem:DeleteFileSystem",

          #          "elasticfilesystem:UntagResource",
          #          "elasticfilesystem:TagResource",
          #
          #            "kms:Encrypt",
          #            "kms:Decrypt",
          #            "kms:ReEncrypt*",
          #            "kms:GenerateDataKey*",
          #            "kms:CreateGrant",
          #            "kms:DescribeKey",
          #            "elasticfilesystem:PutLifecycleConfiguration",
        ]
        "Resource" : "*"
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
