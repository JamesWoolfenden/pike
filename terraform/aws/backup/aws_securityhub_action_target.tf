resource "aws_securityhub_action_target" "pike" {
  depends_on  = [aws_securityhub_account.pike]
  name        = "Send notification"
  identifier  = "SendToChat"
  description = "This is custom action sends selected findings to chat"
}
