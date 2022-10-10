resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "appconfig:DeleteApplication",
          "appconfig:TagResource",
          "appconfig:UpdateApplication",
          "appconfig:UntagResource",
          "appconfig:GetApplication",
          "ec2:DescribeAccountAttributes",
          "appconfig:ListTagsForResource",
          "appconfig:CreateApplication",

          "ec2:DescribeAccountAttributes",
          "appconfig:CreateConfigurationProfile",
          "appconfig:GetConfigurationProfile",
          "appconfig:DeleteConfigurationProfile",
          "appconfig:TagResource",
          "appconfig:UntagResource",
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
