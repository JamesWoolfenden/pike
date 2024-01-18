resource "aws_codecatalyst_project" "pike" {
  provider     = aws.eire
  space_name   = "DeWolf"
  display_name = "Pike"
  description  = "My CodeCatalyst Project created using Terraform"
}
