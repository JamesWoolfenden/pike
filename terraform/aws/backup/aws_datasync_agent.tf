resource "aws_datasync_agent" "pike" {
  ip_address = "1.2.3.4"
  name       = "example"
  tags = {
    pike = "permissions"
  }
}
