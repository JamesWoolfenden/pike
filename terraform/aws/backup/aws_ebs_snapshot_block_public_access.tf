resource "aws_ebs_snapshot_block_public_access" "pike" {
  state = "block-all-sharing"
}
