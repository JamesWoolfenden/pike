resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          //aws_transcribe_language_model
          "transcribe:CreateLanguageModel",
          "transcribe:DescribeLanguageModel",
          "transcribe:DeleteLanguageModel",
          "iam:PassRole",
          "transcribe:TagResource",
          "transcribe:UntagResource",


          //aws_transcribe_medical_vocabulary
          "transcribe:CreateMedicalVocabulary",
          "transcribe:GetMedicalVocabulary",
          "transcribe:DeleteMedicalVocabulary",
          "transcribe:TagResource",
          "transcribe:UntagResource",

          //aws_transcribe_vocabulary
          "transcribe:CreateVocabulary",
          "transcribe:GetVocabulary",
          "transcribe:DeleteVocabulary",
          "transcribe:TagResource",
          "transcribe:UntagResource",

          //aws_transcribe_vocabulary_filter
          "transcribe:CreateVocabularyFilter",
          "transcribe:GetVocabularyFilter",
          "transcribe:UpdateVocabularyFilter",
          "transcribe:ListTagsForResource",
          "transcribe:DeleteVocabularyFilter",
          "transcribe:TagResource",
          "transcribe:UntagResource",

          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
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
          "s3:GetObjectTagging",
          "s3:GetReplicationConfiguration",
          "s3:ListBucket",
          "s3:PutObject"
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
