resource "aws_ram_resource_share" "pike" {
  name                      = "example"
  allow_external_principals = true

  tags = {
    Environment = "Test"
  }
}