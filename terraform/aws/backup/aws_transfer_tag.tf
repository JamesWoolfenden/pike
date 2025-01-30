resource "aws_transfer_tag" "pike" {
  resource_arn = "arn:aws:transfer:us-east-1:680235478471:server/s-01234567890abcdef"
  key          = "aws:transfer:route53HostedZoneId"
  value        = "/hostedzone/MyHostedZoneId"
}
