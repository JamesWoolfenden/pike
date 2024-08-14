resource "aws_chatbot_teams_channel_configuration" "pike" {
  channel_id         = "C07EZ1ABC23"
  configuration_name = "mitt-lags-kanal"
  iam_role_arn       = aws_iam_role.test.arn
  team_id            = "74361522-da01-538d-aa2e-ac7918c6bb92"
  tenant_id          = "1234"

  tags = {
    Name = "mitt-lags-kanal"
  }
}
