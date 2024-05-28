data "aws_chatbot_slack_workspace" "pike" {
  slack_team_name = "pike"
}

output "aws_chatbot_slack_workspace" {
  value = data.aws_chatbot_slack_workspace.pike
}
