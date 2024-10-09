resource "aws_m2_deployment" "pike" {
  environment_id      = "01234567890abcdef012345678"
  application_id      = "34567890abcdef012345678012"
  application_version = 1
  start               = true
}
