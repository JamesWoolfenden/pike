data "aws_sfn_alias" "pike" {
  name="pike"
  statemachine_arn="arn:aws:states:eu-west-2:680235478471:stateMachine:pike/pike"
}