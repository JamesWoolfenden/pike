data "aws_shield_protection" "pike" {
  protection_id = "abc123"
}

output "aws_shield_protection" {
  value = data.aws_shield_protection.pike
}

data "aws_shield_protection" "example" {
  resource_arn = "arn:aws:globalaccelerator::012345678901:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh"
}
