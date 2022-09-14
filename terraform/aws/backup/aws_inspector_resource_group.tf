resource "aws_inspector_resource_group" "pike" {
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
