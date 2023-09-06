data "aws_ec2_spot_price" "pike" {}

output "price" {
  value = data.aws_ec2_spot_price.pike
}
