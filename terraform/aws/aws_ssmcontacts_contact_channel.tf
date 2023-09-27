resource "aws_ssmcontacts_contact_channel" "pike" {
  contact_id = aws_ssmcontacts_contact.pike.arn

  delivery_address {
    simple_address = "email@example.com"
  }

  name = "Example contact channel"
  type = "EMAIL"
}
