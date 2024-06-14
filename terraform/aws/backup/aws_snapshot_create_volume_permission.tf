resource "aws_snapshot_create_volume_permission" "pike" {
  account_id  = "1234567890"
  snapshot_id = aws_ebs_snapshot.example_snapshot.id
}

resource "aws_ebs_volume" "example" {
  availability_zone = "eu-west-2a"
  size              = 40
}

resource "aws_ebs_snapshot" "example_snapshot" {
  volume_id = aws_ebs_volume.example.id
}

#data "aws_caller_identity" "current" {}
