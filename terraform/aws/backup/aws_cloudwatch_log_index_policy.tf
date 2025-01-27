resource "aws_cloudwatch_log_index_policy" "pike" {
  log_group_name = aws_cloudwatch_log_group.test[0].name
  policy_document = jsonencode({
    Fields = ["eventName"]
  })
}
