
resource "aws_cloudwatch_event_bus" "pike" {
  name = "pike"
  tags = {
    pike = "permissions"
  }
}
