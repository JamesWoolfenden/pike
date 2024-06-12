resource "aws_sagemaker_code_repository" "pike" {
  code_repository_name = "example"

  git_config {
    repository_url = "https://github.com/hashicorp/terraform-provider-aws.git"
  }
}
