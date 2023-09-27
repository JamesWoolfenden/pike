resource "aws_account_alternate_contact" "pike" {
  alternate_contact_type = "BILLING"
  email_address          = "someone@gmail.com"
  name                   = "Doctor Suess"
  phone_number           = "+441234567890"
  title                  = "Mr"
}
