data "aws_appstream_image" "pike" {
}

output "aws_appstream_image" {
  value = data.aws_appstream_image.pike
}
