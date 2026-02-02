data "aws_savingsplans_savings_plan" "pike" {
  savings_plan_id = "sp-12345678901234567"
}

output "aws_savingsplans_savings_plan" {
  value = data.aws_savingsplans_savings_plan.pike
}
