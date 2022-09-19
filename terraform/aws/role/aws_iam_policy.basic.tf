resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "config:PutConfigRule",
          "ec2:DescribeAccountAttributes",
          "config:DescribeConfigRules",
          "config:ListTagsForResource",
          "config:DeleteConfigRule",
          "config:TagResource",
          "config:UntagResource",

          "config:PutConfigurationRecorder",
          "config:DescribeConfigurationRecorders",
          "config:DeleteConfigurationRecorder",
          "iam:PassRole",

          "config:DescribeConfigurationRecorderStatus",
          "config:StopConfigurationRecorder",
          "config:StartConfigurationRecorder",

          "config:PutDeliveryChannel",
          "config:DescribeDeliveryChannels",
          "config:DeleteDeliveryChannel",
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
