data "aws_ecr_repositories" "pike" {}

output "repos" {
  value = data.aws_ecr_repositories.pike
}
