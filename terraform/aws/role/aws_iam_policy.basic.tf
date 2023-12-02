resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          //aws_wafregional_geo_match_set
          "waf-regional:GetChangeToken",
          "waf-regional:CreateGeoMatchSet",
          "waf-regional:UpdateGeoMatchSet",
          "waf-regional:GetGeoMatchSet",
          "waf-regional:DeleteGeoMatchSet",



          //aws_wafregional_regex_match_set
          "waf-regional:GetChangeToken",
          "waf-regional:CreateRegexMatchSet",
          "waf-regional:UpdateRegexMatchSet",
          "waf-regional:GetRegexMatchSet",
          "waf-regional:DeleteRegexMatchSet",

          //aws_wafregional_rate_based_rule
          "waf-regional:GetChangeToken",
          "waf-regional:CreateRateBasedRule",
          "waf-regional:UpdateRateBasedRule",
          "waf-regional:GetRateBasedRule",
          "waf-regional:ListTagsForResource",
          "waf-regional:DeleteRateBasedRule",
          "waf-regional:TagResource",
          "waf-regional:UntagResource",

          //aws_wafregional_web_acl
          "waf-regional:GetChangeToken",
          "waf-regional:CreateWebACL",
          "waf-regional:UpdateWebACL",
          "waf-regional:TagResource",
          "waf-regional:GetWebACL",
          "waf-regional:GetLoggingConfiguration",
          "waf-regional:DeleteWebACL",
          "waf-regional:UntagResource",
          "waf-regional:TagResource",

          //aws_wafregional_rule_group
          "waf-regional:GetChangeToken",
          "waf-regional:CreateRuleGroup",
          "waf-regional:UpdateRuleGroup",
          "waf-regional:GetRuleGroup",
          "waf-regional:ListActivatedRulesInRuleGroup",
          "waf-regional:DeleteRuleGroup",
          "waf-regional:UntagResource",
          "waf-regional:TagResource",
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
