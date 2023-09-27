resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_swf_domain
          "swf:RegisterDomain",
          "swf:DescribeDomain",
          "swf:ListTagsForResource",
          "swf:DeprecateDomain",
          "swf:TagResource",

          //aws_ssmcontacts_contact
          "ssm-contacts:CreateContact",
          "ssm-contacts:TagResource",
          //aws_sns_topic_data_protection_policy
          "SNS:PutDataProtectionPolicy",
          "SNS:GetDataProtectionPolicy",
          //aws_sns_sms_preferences
          "SNS:SetSMSAttributes",
          "SNS:GetSMSAttributes"

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
