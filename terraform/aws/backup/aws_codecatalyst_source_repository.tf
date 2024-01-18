resource "aws_codecatalyst_source_repository" "pike" {
  provider     = aws.eire
  name         = "pike"
  project_name = "Pike"
  space_name   = "DeWolf"
}
