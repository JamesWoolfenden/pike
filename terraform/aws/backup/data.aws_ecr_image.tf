data "aws_ecr_image" "pike" {
  repository_name = "pike"
  image_tag       = "pike"
}
