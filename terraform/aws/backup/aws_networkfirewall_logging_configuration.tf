resource "aws_networkfirewall_logging_configuration" "pike" {
  firewall_arn = aws_networkfirewall_firewall.pike.arn
  logging_configuration {
    log_destination_config {
      log_destination = {
        logGroup = "/aws/lambda/message-processor"
      }
      log_destination_type = "CloudWatchLogs"
      log_type             = "ALERT"
    }
  }

}
