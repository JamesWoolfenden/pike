resource "aws_cloudwatch_log_delivery_destination" "example" {
  name = "example"

  delivery_destination_configuration {
    destination_resource_arn = aws_cloudwatch_log_group.test[0].arn
  }

}
