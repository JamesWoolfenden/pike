resource "aws_prometheus_workspace" "example" {
  provider = aws.central
  alias    = "example"
}

resource "aws_cloudwatch_log_group" "example" {
  provider = aws.central
  name     = "/aws/prometheus/query-logs/example"
}

resource "aws_prometheus_query_logging_configuration" "example" {
  provider     = aws.central
  workspace_id = aws_prometheus_workspace.example.id

  destination {
    cloudwatch_logs {
      log_group_arn = "${aws_cloudwatch_log_group.example.arn}:*"
    }

    filters {
      qsp_threshold = 1000
    }
  }

}
