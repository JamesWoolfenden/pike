resource "aws_emr_block_public_access_configuration" "pike" {
  block_public_security_group_rules = true
}
