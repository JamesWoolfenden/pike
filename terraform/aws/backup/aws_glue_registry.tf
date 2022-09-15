resource "aws_glue_registry" "pike" {
  registry_name = "pike-registry"
  description   = "pike description"
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
