data "aws_ebs_snapshot_ids" "pike" {}

output "snaps" {
  value = data.aws_ebs_snapshot_ids.pike
}
