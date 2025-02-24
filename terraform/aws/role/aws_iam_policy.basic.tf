resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          # "cloudwatch:PutInsightRule",
          # "cloudwatch:DescribeInsightRules",
          # "cloudwatch:ListTagsForResource",
          # "cloudwatch:DeleteInsightRules",
          # //,
          #
          # "cloudwatch:TagResource",
          # "cloudwatch:UntagResource"
          "cloudwatch:ListManagedInsightRules",
          "cloudwatch:ListTagsForResource",
          "cloudwatch:DeleteInsightRules",
          "cloudwatch:PutManagedInsightRules"
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
