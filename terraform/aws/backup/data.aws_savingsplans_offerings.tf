data "aws_savingsplans_offerings" "pike" {
}

output "aws_savingsplans_offerings" {
  value = data.aws_savingsplans_offerings.pike
}
