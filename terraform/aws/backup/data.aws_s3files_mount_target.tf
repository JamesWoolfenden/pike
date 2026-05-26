data "aws_s3files_mount_target" "pike" {
  id = "fsmt-01234567890abcdef"
}

output "aws_s3files_mount_target" {
  value = data.aws_s3files_mount_target.pike
}
