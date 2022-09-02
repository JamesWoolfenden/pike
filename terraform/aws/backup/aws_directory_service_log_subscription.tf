resource "aws_directory_service_log_subscription" "pike" {
  directory_id   = aws_directory_service_directory.pike.id
  log_group_name = "pike"
}
