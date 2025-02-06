resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [

          # aws_pinpointsmsvoicev2_configuration_set
          "sms-voice:CreateConfigurationSet",
          "sms-voice:DescribeConfigurationSets",
          "sms-voice:DeleteConfigurationSet",
          "sms-voice:ListTagsForResource",

          # aws_pinpointsmsvoicev2_opt_out_list
          "sms-voice:CreateOptOutList",
          "sms-voice:DescribeOptOutLists",
          "sms-voice:DeleteOptOutList",
          "sms-voice:ListTagsForResource",

          "sms-voice:TagResource",
          "sms-voice:UntagResource",

          # aws_pinpointsmsvoicev2_phone_number
          "sms-voice:ListTagsForResource",
          "sms-voice:DescribePhoneNumbers",
          "sms-voice:RequestPhoneNumber",
          "sms-voice:UpdatePhoneNumber",
          "sms-voice:ReleasePhoneNumber"
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
