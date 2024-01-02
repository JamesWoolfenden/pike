resource "aws_verifiedaccess_instance_logging_configuration" "pike" {
  access_logs {
    cloudwatch_logs {
      enabled   = true
      log_group = aws_cloudwatch_log_group.pike.id
    }
  }
  verifiedaccess_instance_id = aws_verifiedaccess_instance.pike.id
}


resource "aws_cloudwatch_log_group" "pike" {
}
