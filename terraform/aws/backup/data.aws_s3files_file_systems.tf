data "aws_s3files_file_systems" "pike" {
}

output "aws_s3files_file_systems" {
  value = data.aws_s3files_file_systems.pike
}
