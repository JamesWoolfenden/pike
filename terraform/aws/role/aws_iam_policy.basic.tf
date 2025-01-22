resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          # aws_s3tables_namespace
          "s3tables:CreateNamespace",
          "s3tables:GetNamespace",
          "s3tables:DeleteNamespace",

          # aws_s3tables_table
          "s3tables:GetTable",
          "s3tables:CreateTable",
          "s3tables:DeleteTable",

          # aws_s3tables_policy
          "s3tables:GetTablePolicy",
          "s3tables:DeleteTablePolicy",
          "s3tables:PutTablePolicy",

          # others
          "s3tables:CreateTableBucket",
          "s3tables:PutTableBucketMaintenanceConfiguration",
          "s3tables:GetTableBucket",
          "s3tables:GetTableBucketMaintenanceConfiguration",
          "s3tables:DeleteTableBucket"
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
