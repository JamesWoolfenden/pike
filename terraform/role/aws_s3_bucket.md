apiresource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:GetLifecycleConfiguration",
          "s3:GetBucketTagging",
          "s3:GetBucketWebsite",
          "s3:GetBucketLogging",
          "s3:CreateBucket",
          "s3:ListBucket",
          "s3:GetAccelerateConfiguration",
          "s3:GetBucketVersioning",
          "s3:GetBucketAcl",
          "s3:GetBucketPolicy",
          "s3:GetReplicationConfiguration",
          "s3:GetBucketObjectLockConfiguration",
          "s3:GetObjectAcl",
          "s3:GetObject",
          "s3:GetEncryptionConfiguration",
          "s3:GetBucketRequestPayment",
          "s3:GetBucketCORS",
          "s3:DeleteBucket",

          "s3:GetBucketObjectLockConfiguration",
          "s3:GetObjectLegalHold",
          "s3:GetObjectRetention",
          "s3:PutBucketObjectLockConfiguration",
          "s3:PutObjectLegalHold",
          "s3:PutObjectRetention",
          "s3:PutObject",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
