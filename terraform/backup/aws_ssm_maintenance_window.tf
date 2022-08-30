resource "aws_ssm_maintenance_window" "production" {
  name     = "maintenance-window-application"
  schedule = "cron(0 16 ? * TUE *)"
  duration = 5
  cutoff   = 1
  tags = {
    pike   = "Permissions"
    delete = "me"
  }
}
