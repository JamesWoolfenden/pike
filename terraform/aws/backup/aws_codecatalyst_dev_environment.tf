resource "aws_codecatalyst_dev_environment" "pike" {
  provider      = aws.eire
  alias         = "devenv"
  space_name    = "DeWolf"
  project_name  = "myproject"
  instance_type = "dev.standard1.small"

  persistent_storage {
    size = 16
  }

  ides {
    name    = "PyCharm"
    runtime = "public.ecr.aws/jetbrains/py"
  }

  inactivity_timeout_minutes = 40

  repositories {
    repository_name = "terraform-provider-aws"
    branch_name     = "main"
  }

}
