resource "aws_datazone_asset_type" "pike" {
  description               = "example"
  domain_identifier         = aws_datazone_domain.test.id
  name                      = "example"
  owning_project_identifier = aws_datazone_project.test.id
}
