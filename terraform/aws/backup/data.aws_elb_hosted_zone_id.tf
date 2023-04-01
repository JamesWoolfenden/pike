data "aws_elb_hosted_zone_id" "pike" {}

output "zone" {

  value = data.aws_elb_hosted_zone_id.pike
}
