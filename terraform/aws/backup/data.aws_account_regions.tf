data "aws_account_regions" "pike" {
}

output "aws_account_regions" {
  value = data.aws_account_regions.pike
}
