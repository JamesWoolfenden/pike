resource "aws_ssmcontacts_contact" "pike" {
  alias        = "jameswoolfenden"
  type         = "PERSONAL"
  display_name = "james woolfenden"
  tags = {
    pike = "permissions"
  }
}
