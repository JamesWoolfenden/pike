resource "aws_ec2_availability_zone_group" "pike" {
  group_name    = "us-west-2-lax-1"
  opt_in_status = "opted-in"
}