resource "aws_servicequotas_template" "pike" {
  provider     = aws.central
  quota_code   = "L-2ACBD22F" # function and layer storage (default: 75 GB)
  service_code = "lambda"
  value        = "80"
  region       = "eu-west-2"
}
