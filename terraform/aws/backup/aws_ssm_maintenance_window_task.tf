resource "aws_ssm_maintenance_window_task" "example" {
  max_concurrency = 2
  max_errors      = 1
  priority        = 1
  task_arn        = "AWS-RestartEC2Instance"
  task_type       = "AUTOMATION"
  window_id       = aws_ssm_maintenance_window.production.id

  targets {
    key    = "InstanceIds"
    values = [aws_instance.example.id]
  }

  task_invocation_parameters {
    automation_parameters {
      document_version = "$LATEST"

      parameter {
        name   = "InstanceId"
        values = [aws_instance.example.id]
      }
    }
  }
}
