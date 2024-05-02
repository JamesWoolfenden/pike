resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          //backup
          "fsx:CreateBackup",
          "fsx:DeleteBackup",
          "fsx:TagResource",
          "fsx:UntagResource",

          //aws_fsx_lustre_file_system & aws_fsx_windows_file_system
          "fsx:DescribeFileSystems",
          "fsx:CreateFileSystem",
          "fsx:DeleteFileSystem",
          "fsx:UpdateFileSystem",
          "fsx:TagResource",
          "fsx:UntagResource",

          //cache
          "fsx:CreateFileCache",
          "fsx:DeleteFileCache",
          "fsx:UpdateFileCache",
          "ec2:DescribeVpcs",
          "fsx:DescribeFileCaches",
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "fsx:ListTagsForResource",
          "fsx:TagResource",
          "fsx:UntagResource",

          //aws_fsx_data_repository_association
          "fsx:CreateDataRepositoryAssociation",
          "fsx:DeleteDataRepositoryAssociation",
          "fsx:UpdateDataRepositoryAssociation",
          "fsx:DescribeDataRepositoryAssociations",
          "fsx:TagResource",
          "fsx:UntagResource",
          "iam:CreateServiceLinkedRole",
          "iam:AttachRolePolicy",
          "iam:PutRolePolicy",

          //lustre
          "iam:CreateServiceLinkedRole",
          "iam:AttachRolePolicy",
          "iam:PutRolePolicy",
          "fsx:TagResource",
          "fsx:UntagResource",

          //storage
          "fsx:CreateStorageVirtualMachine",
          "fsx:DeleteStorageVirtualMachine",
          "fsx:UpdateStorageVirtualMachine",
          "fsx:TagResource",
          "fsx:UntagResource",

          //volume
          "fsx:CreateVolume",
          "fsx:DeleteVolume",
          "fsx:UpdateVolume",
          "fsx:TagResource",
          "fsx:UntagResource",

          "ec2:DescribeSubnets",
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:CreateSecurityGroup",
          "ec2:CreateTags",
          "ec2:DeleteSecurityGroup",
          "ec2:DeleteTags",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeImages",
          "ec2:DescribeInstanceAttribute",
          "ec2:DescribeInstanceCreditSpecifications",
          "ec2:DescribeInstanceTypes",
          "ec2:DescribeInstances",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeTags",
          "ec2:DescribeVolumes",
          "ec2:ModifyInstanceAttribute",
          "ec2:RevokeSecurityGroupEgress",
          "ec2:RunInstances",
          "ec2:StartInstances",
          "ec2:StopInstances",
          "ec2:TerminateInstances",
          "iam:CreateRole",
          "iam:DeleteRole",
          "iam:DeleteRolePolicy",
          "iam:GetRole",
          "iam:GetRolePolicy",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",
          "iam:PutRolePolicy",
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
          "s3:PutBucketPolicy",
          "s3:PutObject",
          "s3:PutBucketAcl",
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
