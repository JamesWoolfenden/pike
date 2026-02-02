data "aws_arcregionswitch_plan" "pike" {
  arn = "arn:aws:arcregionswitch:us-west-2:123456789012:plan/example-plan"
}

output "aws_arcregionswitch_plan" {
  value = data.aws_arcregionswitch_plan.pike
}
