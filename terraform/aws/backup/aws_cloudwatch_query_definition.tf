resource "aws_cloudwatch_query_definition" "pike" {
  name = "custom_query"

  log_group_names = [
    "/aws/logGroup1",
    "/aws/logGroup2"
  ]

  query_string = <<EOF
fields @timestamp, @message
| sort @timestamp desc
| limit 25
EOF
}

output "definition" {
  value = aws_cloudwatch_query_definition.pike
}
