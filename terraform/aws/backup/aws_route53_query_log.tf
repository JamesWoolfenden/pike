resource "aws_route53_query_log" "pike" {
  cloudwatch_log_group_arn = aws_cloudwatch_log_group.querylog.arn
  zone_id                  = "Z0613304D03LG1SU5BI"
}


provider "aws" {
  profile = "basic"
  alias   = "us-east-1"
  region  = "us-east-1"
}

resource "aws_cloudwatch_log_group" "querylog" {
  provider = aws.us-east-1

  name              = "querylog"
  retention_in_days = 30
}

# Example CloudWatch log resource policy to allow Route53 to write logs
# to any log group under /aws/route53/*

data "aws_iam_policy_document" "route53-query-logging-policy" {
  statement {
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = ["arn:aws:logs:*"]

    principals {
      identifiers = ["route53.amazonaws.com"]
      type        = "Service"
    }
  }
}

resource "aws_cloudwatch_log_resource_policy" "route53-query-logging-policy" {
  provider = aws.us-east-1

  policy_document = data.aws_iam_policy_document.route53-query-logging-policy.json
  policy_name     = "route53-query-logging-policy"
}

#aws route53 create-query-logging-config --hosted-zone-id Z0613304D03LG1SU5BI --cloud-watch-logs-log-group-arn arn:aws:logs:us-east-1:680235478471:log-group:querylog --region us-east-1
