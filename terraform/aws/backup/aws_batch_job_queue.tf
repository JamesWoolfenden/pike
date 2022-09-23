resource "aws_batch_job_queue" "pike" {
  name     = "test"
  state    = "ENABLED"
  priority = 1
  compute_environments = [
    aws_batch_compute_environment.pike.arn,
  ]
  scheduling_policy_arn = aws_batch_scheduling_policy.pike.arn
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
