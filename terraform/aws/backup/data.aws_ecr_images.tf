data "aws_ecr_images" "pike" {
  repository_name = "pike"
}

output "aws_ecr_images" {
  value = data.aws_ecr_images.pike
}
