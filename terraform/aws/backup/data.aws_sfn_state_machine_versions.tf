data "aws_sfn_state_machine_versions" "pike" {
  statemachine_arn = data.aws_sfn_state_machine.pike.arn
}

output "version" {
  value = data.aws_sfn_state_machine_versions.pike
}
