data "aws_s3files_file_system" "pike" {
  id = "fs-01234567890abcdef"
}

output "aws_s3files_file_system" {
  value = data.aws_s3files_file_system.pike
}
