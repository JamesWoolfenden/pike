data "aws_route53_records" "pike" {
  zone_id = "Z0613304D03LG1SU5BI"
}

output "aws_route53_records" {
  value = data.aws_route53_records.pike
}
