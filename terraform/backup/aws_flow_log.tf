resource "aws_flow_log" "test" {
  iam_role_arn    = "arn:aws:iam::680235478471:role/test-flowlog"
  log_destination = "arn:aws:logs:eu-west-2:680235478471:log-group:pike-flow-log"
  traffic_type    = "ALL"
  vpc_id          = "vpc-0c33dc8cd64f408c4"
  tags = {
    pike = "permissions"
  }
}
