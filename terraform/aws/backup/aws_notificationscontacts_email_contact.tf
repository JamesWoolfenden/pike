resource "aws_notificationscontacts_email_contact" "pike" {
  name          = "pike-contact2"
  email_address = "pike3@pike.com"

  tags = {
    pike        = "permissions"
    Environment = "Production"
  }
}
