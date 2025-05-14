resource "aws_ec2_default_credit_specification" "pike" {
  instance_family = "t2"
  cpu_credits     = "standard"
}
