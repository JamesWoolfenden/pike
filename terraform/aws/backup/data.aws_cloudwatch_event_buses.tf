data "aws_cloudwatch_event_buses" "pike" {
  name_prefix = "pike"
}

output "aws_cloudwatch_event_buses" {
  value = data.aws_cloudwatch_event_buses.pike
}
