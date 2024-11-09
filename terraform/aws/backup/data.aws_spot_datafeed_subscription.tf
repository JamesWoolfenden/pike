data "aws_spot_datafeed_subscription" "pike" {
}

output "aws_spot_datafeed_subscription" {
  value = data.aws_spot_datafeed_subscription.pike
}
