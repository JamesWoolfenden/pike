resource "aws_codecommit_repository" "test" {
  repository_name = "test"
}

resource "aws_codecommit_trigger" "test" {
  repository_name = aws_codecommit_repository.test.repository_name

  trigger {
    name            = "all"
    events          = ["all"]
    destination_arn = "arn:aws:sns:eu-west-2:680235478471:My_Alert"
  }

}
