resource "aws_dms_replication_task" "pike" {
  cdc_start_time           = "1993-05-21T05:50:00Z"
  migration_type           = "full-load"
  replication_instance_arn = aws_dms_replication_instance.pike.replication_instance_arn
  replication_task_id      = "test-dms-replication-task-tf"
  source_endpoint_arn      = aws_dms_endpoint.pike.endpoint_arn
  table_mappings           = "{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}"

  tags = {
    Name = "test"
  }

  target_endpoint_arn = aws_dms_endpoint.pike.endpoint_arn
}
