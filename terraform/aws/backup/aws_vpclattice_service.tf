resource "aws_vpclattice_service" "pike" {
  name               = "example"
  auth_type          = "AWS_IAM"
  custom_domain_name = "example.com"

}
