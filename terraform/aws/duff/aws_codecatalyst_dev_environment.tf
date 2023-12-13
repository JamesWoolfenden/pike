#resource "aws_codecatalyst_dev_environment" "pike" {
#  alias         = "devenv"
#  space_name    = "myspace"
#  project_name  = "myproject"
#  instance_type = "dev.standard1.small"
#
#  persistent_storage {
#    size =64
#  }
#
#  ides {
#    name    = "PyCharm"
#    runtime = "public.ecr.aws/jetbrains/py"
#  }
#
#  inactivity_timeout_minutes = 40
#
#  repositories {
#    repository_name = "terraform-provider-aws"
#    branch_name     = "main"
#  }
#}
