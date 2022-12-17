resource "aws_ebs_snapshot_copy" "pike" {
  source_region      = "eu-west-1"
  source_snapshot_id = "snap-08f14c31fd87f4ab7"
  tags = {
    pike = "permission"
  }
}
