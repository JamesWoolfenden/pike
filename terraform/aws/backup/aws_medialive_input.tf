resource "aws_medialive_input" "pike" {
  name                  = "example-input"
  input_security_groups = [aws_medialive_input_security_group.pike.id]
  type                  = "UDP_PUSH"

  tags = {
    pike = "permissions"
  }
}
