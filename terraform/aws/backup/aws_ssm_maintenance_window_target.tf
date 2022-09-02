resource "aws_ssm_maintenance_window_target" "target1" {
  window_id     = aws_ssm_maintenance_window.production.id
  name          = "maintenance-window-target"
  description   = "This is a maintenance window target update"
  resource_type = "INSTANCE"

  targets {
    key    = "tag:Name"
    values = ["acceptance_test"]
  }

}
