resource "aws_ssmcontacts_plan" "pike" {
  contact_id = aws_ssmcontacts_contact.pike.arn
  stage {
    duration_in_minutes = 2
  }
}
