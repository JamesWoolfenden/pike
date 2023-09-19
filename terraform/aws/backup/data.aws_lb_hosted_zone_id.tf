data "aws_lb_hosted_zone_id" "pike" {}

output "zone_id" {
  value = data.aws_lb_hosted_zone_id.pike
}
