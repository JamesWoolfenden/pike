resource "aws_iam_instance_profile" "example" {
  name = "test_profile"
  role = "lambda_basic"
  #  tags = {
  #    test="james"
  #  }
}
