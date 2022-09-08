resource "aws_appautoscaling_scheduled_action" "dynamodb" {
  name               = "dynamodb"
  service_namespace  = aws_appautoscaling_target.dynamodb_table_read_target.service_namespace
  resource_id        = aws_appautoscaling_target.dynamodb_table_read_target.resource_id
  scalable_dimension = aws_appautoscaling_target.dynamodb_table_read_target.scalable_dimension
  schedule           = "at(2006-01-02T15:04:05)"

  scalable_target_action {
    min_capacity = 1
    max_capacity = 200
  }
}
