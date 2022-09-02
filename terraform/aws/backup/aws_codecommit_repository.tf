resource "aws_codecommit_repository" "example" {
  repository_name = "MyNewTestRepository"
  description     = "This is the Sample App Repository again"
  default_branch  = "masterblaster"
  tags = {
    pike = "permissions"
    test = "update"
  }
}
