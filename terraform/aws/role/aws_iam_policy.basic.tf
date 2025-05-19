resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          "cloudfront:CreateKeyValueStore",
          "cloudfront:DeleteKeyValueStore",
          "cloudfront:DescribeKeyValueStore",
          "cloudfront:UpdateKeyValueStore",

          //aws_bedrockagent_prompt
          "bedrock:CreatePrompt",
          "bedrock:DeletePrompt",
          "bedrock:UpdatePrompt",
          "bedrock:GetPrompt",
          "bedrock:ListTagsForResource",
          "bedrock:UntagResource",
          "bedrock:TagResource",
          "kms:GenerateDataKey",
          "kms:Decrypt",

          //aws_cloudfrontkeyvaluestore_keys_exclusive
          "cloudfront-keyvaluestore:DescribeKeyValueStore",
          "cloudfront-keyvaluestore:ListKeys",
          "cloudfront-keyvaluestore:UpdateKeys",

          //aws_dataexchange_revision_assets
          "dataexchange:TagResource",
          "dataexchange:ListTagsForResource",
          "dataexchange:UntagResource",
          "dataexchange:CreateRevision",
          "dataexchange:DeleteRevision",
          "dataexchange:UpdateRevision",
          "dataexchange:GetRevision",

          //aws_inspector2_filter
          "inspector2:TagResource",
          "inspector2:ListTagsForResource",
          "inspector2:UntagResource",
          "inspector2:CreateFilter",
          "inspector2:ListFilters",
          "inspector2:DeleteFilter",
          "inspector2:UpdateFilter",

          //aws_wafv2_api_key
          "wafv2:CreateAPIKey",
          "wafv2:DeleteAPIKey",
          "wafv2:ListAPIKeys",

          //aws_account_primary_contact
          "account:GetContactInformation",

          //aws_dynamodb_tables
          "dynamodb:ListTables"
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
