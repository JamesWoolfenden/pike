data "aws_cloudwatch_event_source" "pike" {
  name_prefix = "aws.partner/examplepartner.com"
}
