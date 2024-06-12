resource "aws_sagemaker_monitoring_schedule" "pike" {
  name = "my-monitoring-schedule"
  monitoring_schedule_config {
    monitoring_job_definition_name = aws_sagemaker_data_quality_job_definition.pike.name
    monitoring_type                = "DataQuality"
  }
}
