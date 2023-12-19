resource "aws_codebuild_webhook" "pike" {
  project_name = aws_codebuild_project.example.name
  build_type   = "BUILD"
  filter_group {
    filter {
      type    = "EVENT"
      pattern = "PULL_REQUEST_CREATED"
    }

    #    filter {
    #      type    = "BASE_REF"
    #      pattern = "main"
    #    }
  }

}


#resource "aws_codebuild_project" "example" {
#  name         = "pike"
#  service_role = "arn:aws:iam::680235478471:role/codebuild"
#  artifacts {
#    type = "NO_ARTIFACTS"
#  }
#  environment {
#    compute_type                = "BUILD_GENERAL1_SMALL"
#    image                       = "aws/codebuild/amazonlinux2-x86_64-standard:4.0"
#    type                        = "LINUX_CONTAINER"
#    image_pull_credentials_type = "CODEBUILD"
#
#    environment_variable {
#      name  = "SOME_KEY1"
#      value = "SOME_VALUE1"
#    }
#
#    environment_variable {
#      name  = "SOME_KEY2"
#      value = "SOME_VALUE2"
#      type  = "PARAMETER_STORE"
#    }
#  }
#  source {
#    type            = "GITHUB"
#    location        = "https://github.com/mitchellh/packer.git"
#    git_clone_depth = 1
#
#    git_submodules_config {
#      fetch_submodules = true
#    }
#  }
#}
