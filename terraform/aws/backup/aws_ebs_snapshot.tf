resource "aws_ebs_snapshot" "pike" {
  volume_id    = data.aws_ebs_volume.pike.id
  description  = "pike"
  storage_tier = "standard"

  tags = {
    pike = "permissions"
    and  = "another"
  }
}
