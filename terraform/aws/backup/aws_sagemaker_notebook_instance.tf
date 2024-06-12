resource "aws_sagemaker_notebook_instance" "pike" {
  name          = "my-notebook-instance"
  role_arn      = aws_iam_role.example.arn
  instance_type = "ml.t2.medium"

  tags = {
    pike = "permissions"
    Name = "foo"
  }
}
