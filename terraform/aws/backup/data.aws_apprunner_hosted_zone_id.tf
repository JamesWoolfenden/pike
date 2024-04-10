data "aws_apprunner_hosted_zone_id" "pike" {
}

output "aws_apprunner_hosted_zone_id" {
  value = data.aws_apprunner_hosted_zone_id.pike
}
