resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_storagegateway_cache not implemented
          "storagegateway:DescribeCache",
          "storagegateway:AddCache",

          //aws_storagegateway_cached_iscsi_volume not implemented
          "storagegateway:DescribeCachediSCSIVolumes",
          "storagegateway:CreateCachediSCSIVolume",

          //aws_storagegateway_file_system_association not implemented
          "storagegateway:DescribeFileSystemAssociations",
          "storagegateway:UpdateFileSystemAssociation",
          "storagegateway:RemoveTagsFromResource",
          "storagegateway:AddTagsToResource",
          "storagegateway:ListTagsForResource",

          //aws_fsx_windows_file_system not implemented
          "fsx:DeleteFileSystem",
          "fsx:CreateFileSystem",
          "fsx:DescribeFileSystems",
          "fsx:UpdateFileSystem",

          //aws_storagegateway_gateway not implemented
          "storagegateway:DescribeGatewayInformation",
          "storagegateway:ActivateGateway",
          "storagegateway:DeleteGateway",
          "storagegateway:DisableGateway",
          "storagegateway:RemoveTagsFromResource",
          "storagegateway:AddTagsToResource",
          "storagegateway:ListTagsForResource",

          //aws_storagegateway_nfs_file_share not implemented
          "storagegateway:DescribeNFSFileShares",
          "storagegateway:CreateNFSFileShare",
          "storagegateway:UpdateNFSFileShare",
          "storagegateway:RemoveTagsFromResource",
          "storagegateway:AddTagsToResource",
          "storagegateway:ListTagsForResource",

          //aws_storagegateway_smb_file_share not implemented
          "storagegateway:DescribeSMBFileShares",
          "storagegateway:UpdateSMBFileShare",
          "storagegateway:CreateSMBFileShare",
          "storagegateway:RemoveTagsFromResource",
          "storagegateway:AddTagsToResource",
          "storagegateway:ListTagsForResource",

          //aws_storagegateway_stored_iscsi_volume not implemented
          "storagegateway:DescribeStorediSCSIVolumes",
          "storagegateway:CreateStorediSCSIVolume",
          "storagegateway:RemoveTagsFromResource",
          "storagegateway:AddTagsToResource",
          "storagegateway:ListTagsForResource",

          //aws_storagegateway_tape_pool not implemented
          "storagegateway:ListTapePools",
          "storagegateway:CreateTapePool",
          "storagegateway:DeleteTapePool",
          "storagegateway:RemoveTagsFromResource",
          "storagegateway:AddTagsToResource",
          "storagegateway:ListTagsForResource",

          //aws_storagegateway_upload_buffer not implemented
          "storagegateway:DescribeUploadBuffer",
          "storagegateway:AddUploadBuffer",

          //aws_storagegateway_working_storage not implemented
          "storagegateway:AddWorkingStorage",
          "storagegateway:DescribeWorkingStorage",

          "ec2:CreateTags",
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:AttachVolume",
          "ec2:CreateVolume",
          "ec2:DeleteVolume",
          "ec2:DescribeSubnets",
          "ec2:DescribeImages",
          "ec2:DescribeInstanceAttribute",
          "ec2:DescribeInstanceCreditSpecifications",
          "ec2:DescribeInstanceTypes",
          "ec2:DescribeInstances",
          "ec2:DescribeTags",
          "ec2:DescribeVolumes",
          "ec2:DetachVolume",
          "ec2:ModifyInstanceAttribute",
          "ec2:ModifyVolume",
          "ec2:RunInstances",
          "ec2:StartInstances",
          "ec2:StopInstances",
          "ec2:TerminateInstances",
          "iam:CreateRole",
          "iam:DeleteRole",
          "iam:GetRole",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",
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
          "s3:PutObject",
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
