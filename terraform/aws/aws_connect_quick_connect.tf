resource "aws_connect_quick_connect" "pike" {

  instance_id = aws_connect_instance.pike.id
  name        = "Example Name"
  description = "quick connect phone number"

  quick_connect_config {
    quick_connect_type = "PHONE_NUMBER"

    phone_config {
      phone_number = "+12345678912"
    }
  }
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
