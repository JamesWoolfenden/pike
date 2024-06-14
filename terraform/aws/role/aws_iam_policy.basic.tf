resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:CreateSnapshot",
          "ec2:CreateVolume",
          "ec2:DeleteSnapshot",
          "ec2:DeleteVolume",
          "ec2:DescribeSnapshots",
          "ec2:DescribeVolumes",
          "ec2:ModifyVolume",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",

          //aws_sns_platform_application
          "sns:CreatePlatformApplication",
          "sns:DeletePlatformApplication",
          "sns:SetPlatformApplicationAttributes",
          "sns:GetPlatformApplicationAttributes",

          //aws_simpledb_domain
          "sdb:CreateDomain",
          "sdb:DomainMetadata",
          "sdb:DeleteDomain",

          //aws_snapshot_create_volume_permission
          "ec2:ModifySnapshotAttribute"
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
