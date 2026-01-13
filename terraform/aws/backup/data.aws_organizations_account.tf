data "aws_organizations_account" "pike" {
  account_id = "pike-1234567"
}

output "aws_organizations_account" {
  value = data.aws_organizations_account.pike
}
