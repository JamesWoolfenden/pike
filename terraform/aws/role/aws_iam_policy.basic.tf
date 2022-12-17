resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:DescribeAccountAttributes",
          "kms:DescribeKey",

          #config
          "xray:PutEncryptionConfig",
          "xray:GetEncryptionConfig",
          "kms:CreateGrant",
          "kms:DescribeKey",

          #group
          "xray:CreateGroup",
          "xray:ListTagsForResource",
          "xray:GetGroup",
          "xray:DeleteGroup",
          "xray:UntagResource",
          "xray:TagResource",

          #sampling rule
          "xray:CreateSamplingRule",
          "xray:GetSamplingRules",
          "xray:ListTagsForResource",
          "xray:DeleteSamplingRule",
          "xray:UntagResource",
          "xray:TagResource",
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
