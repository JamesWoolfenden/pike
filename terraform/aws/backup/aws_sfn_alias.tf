resource "aws_sfn_alias" "pike" {
  name = "pike"
  routing_configuration {
    state_machine_version_arn = "arn:aws:states:eu-west-2:680235478471:stateMachine:MyStateMachine-ltmqi7y6q:1"
    weight                    = 100
  }
}
