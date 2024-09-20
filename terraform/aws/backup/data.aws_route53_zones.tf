data "aws_route53_zones" "pike" {
}

output "aws_route53_zones" {
  value = data.aws_route53_zones.pike
}
