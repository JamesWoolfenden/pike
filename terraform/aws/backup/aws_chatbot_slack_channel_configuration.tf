resource "aws_chatbot_slack_channel_configuration" "pike" {
  configuration_name = "min-slaka-kanal"
  iam_role_arn       = aws_iam_role.test.arn
  slack_channel_id   = "C07EZ1ABC23"
  slack_team_id      = "T07EA123LEP"

  tags = {
    Name = "min-slaka-kanal"
  }
}
