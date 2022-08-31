data "aws_cloudwatch_log_group" "pike" {
  name = "/aws/lambda/reads3"
}

output "logs" {
  value = data.aws_cloudwatch_log_group.pike
}
