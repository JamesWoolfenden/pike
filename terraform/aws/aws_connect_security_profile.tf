resource "aws_connect_security_profile" "pike" {
  instance_id = aws_connect_instance.pike.id
  name        = "example"
  description = "example description"

  permissions = [
    "BasicAgentAccess",
    "OutboundCallAccess",
  ]
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
