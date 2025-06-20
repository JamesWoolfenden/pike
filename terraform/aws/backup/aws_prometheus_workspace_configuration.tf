resource "aws_prometheus_workspace_configuration" "pike" {
  workspace_id             = aws_prometheus_workspace.example.id
  retention_period_in_days = 60

  limits_per_label_set {
    label_set = {
      "env" = "dev"
    }
    limits {
      max_series = 100000
    }
  }

  limits_per_label_set {
    label_set = {
      "env" = "prod"
    }
    limits {
      max_series = 400000
    }
  }
}

resource "aws_prometheus_workspace" "example" {

}
