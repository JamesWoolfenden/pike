resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_oam_sinks
          "oam:ListSinks",
          //aws_oam_sink
          "oam:GetSink",
          //aws_oam_links
          "oam:ListLinks",
          //aws_oam_link
          "oam:GetLink",
          //aws_lex_slot_type
          "lex:GetSlotType",
          //aws_lex_intent
          "lex:GetIntent",
          //aws_lex_bot_alias
          "lex:GetBotAlias",
          //aws_lex_bot
          "lex:GetBot",
          //aws_kendra_thesaurus
          "kendra:DescribeThesaurus",
          //aws_kendra_query_suggestions_block_list
          "kendra:DescribeQuerySuggestionsBlockList",
          //aws_kendra_index
          "kendra:DescribeIndex",
          //aws_kendra_faq
          "kendra:DescribeFaq",
          //aws_kendra_experience
          "kendra:DescribeExperience"
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
