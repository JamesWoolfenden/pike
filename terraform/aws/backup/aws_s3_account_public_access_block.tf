resource "aws_s3_account_public_access_block" "pike" {
  block_public_acls   = true
  block_public_policy = true
}
