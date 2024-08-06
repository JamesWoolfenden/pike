data "aws_ecr_repository_creation_template" "pike" {
  prefix = "pike"
}

output "aws_ecr_repository_creation_template" {
  value = data.aws_ecr_repository_creation_template.pike
}
