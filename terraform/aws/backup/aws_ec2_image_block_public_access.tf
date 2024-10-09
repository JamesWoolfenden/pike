resource "aws_ec2_image_block_public_access" "pike" {
  state = "block-new-sharing"
}
