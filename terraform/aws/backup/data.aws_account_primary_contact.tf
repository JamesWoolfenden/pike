data "aws_account_primary_contact" "pike" {
}

output "aws_account_primary_contact" {
  value = data.aws_account_primary_contact.pike
}
