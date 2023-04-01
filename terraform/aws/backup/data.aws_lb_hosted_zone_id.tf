data "aws_lb_hosted_zone_id" "pike" {}

output "id" {
  value = data.aws_lb_hosted_zone_id.pike
}
