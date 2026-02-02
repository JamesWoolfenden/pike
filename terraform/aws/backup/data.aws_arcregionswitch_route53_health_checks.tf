data "aws_arcregionswitch_route53_health_checks" "pike" {
  plan_arn = "arn:aws:arc-region-switch::123456789012:plan/example-plan:abc123"
}

output "aws_arcregionswitch_route53_health_checks" {
  value = data.aws_arcregionswitch_route53_health_checks.pike
}
