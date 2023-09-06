data "aws_s3_objects" "pike" {
  bucket = "pike-680235478471"
}

output "objects" {
  value = data.aws_s3_objects.pike
}
