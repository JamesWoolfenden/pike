data "aws_ecrpublic_images" "pike" {
  provider        = aws.central
  repository_name = "pike"
}

output "aws_ecrpublic_images" {
  value = data.aws_ecrpublic_images.pike
}
