resource "aws_m2_environment" "pike" {
  name          = "test-env"
  engine_type   = "bluage"
  instance_type = "M2.m5.large"
  subnet_ids    = ["subnet-01234567890abcdef", "subnet-01234567890abcdea"]
}
