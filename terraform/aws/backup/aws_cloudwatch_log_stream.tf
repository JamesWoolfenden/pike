resource "aws_cloudwatch_log_stream" "pike" {
  log_group_name = "pike"
  name           = "pike"
}
