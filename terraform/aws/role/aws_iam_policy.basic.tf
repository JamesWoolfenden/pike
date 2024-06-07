resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_ses_template
          "ses:CreateTemplate",
          "ses:GetTemplate",
          "ses:DeleteTemplate",
          "ses:UpdateTemplate",

          //aws_ses_receipt_filter
          "ses:CreateReceiptFilter",
          "ses:ListReceiptFilters",
          "ses:DeleteReceiptFilter",

          //aws_ses_email_identity
          "ses:VerifyEmailIdentity",
          "ses:GetIdentityVerificationAttributes",
          "ses:DeleteIdentity",

          //aws_ses_configuration_set
          "ses:CreateConfigurationSet",
          "ses:PutConfigurationSetDeliveryOptions",
          "ses:DescribeConfigurationSet",
          "ses:DeleteConfigurationSet",
          "ses:CreateConfigurationSetTrackingOptions",

          //aws_ses_active_receipt_rule_set
          "ses:SetActiveReceiptRuleSet",
          "ses:ListReceiptRuleSets",
          "ses:CreateReceiptRuleSet",
          "ses:DescribeReceiptRuleSet",
          "ses:DeleteReceiptRuleSet",
          "ses:DescribeActiveReceiptRuleSet",

          //aws_ses_event_destination
          "ses:CreateConfigurationSetEventDestination",
          "ses:DeleteConfigurationSetEventDestination",

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
