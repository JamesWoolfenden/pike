resource "aws_secretsmanager_secret_rotation" "pike" {
  rotation_lambda_arn = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
  secret_id           = data.aws_secretsmanager_secret.pike.id
  rotation_rules {
    automatically_after_days = 7
  }

}
