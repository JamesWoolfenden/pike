data "aws_ebs_snapshot" "pike" {
  most_recent = true
}

output "snap" {
  value = data.aws_ebs_snapshot.pike
}
