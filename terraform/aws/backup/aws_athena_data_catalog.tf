resource "aws_athena_data_catalog" "pike" {
  name        = "athena-data-catalog"
  description = "Example Athena data catalog"
  type        = "LAMBDA"

  parameters = {
    "function" = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
  }

  tags = {
    Pike = "Permissions"
  }
}
