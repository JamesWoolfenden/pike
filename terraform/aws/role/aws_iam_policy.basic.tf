resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [


          //aws_cloudwatch_event_bus
          "events:TagResource",
          "events:DescribeEventBus",
          "events:ListTagsForResource",
          "events:DeleteEventBus",
          "events:CreateEventBus",


          //healthcheck
          "route53:ChangeTagsForResource",
          "route53:CreateHealthCheck",
          "route53:DeleteHealthCheck",
          "route53:GetHealthCheck",
          "route53:GetHealthCheckCount",
          "route53:GetHealthCheckLastFailureReason",
          "route53:GetHealthCheckStatus",
          "route53:ListHealthChecks",
          "route53:ListTagsForResource",
          "route53:ListTagsForResources",
          "route53:UpdateHealthCheck",

          //aws_cloudwatch_log_destination
          "logs:PutDestination",
          "iam:PassRole",
          "logs:TagResource",
          "logs:UntagResource",
          "logs:ListTagsForResource",
          "logs:DescribeDestinations",
          "logs:DeleteDestination",

          //aws_cloudwatch_log_destination_policy
          "logs:PutDestinationPolicy",
          "logs:GetDeliveryDestinationPolicy",

          //aws_kinesis_stream_consumer
          "kinesis:RegisterStreamConsumer",
          "kinesis:DescribeStreamConsumer",
          "kinesis:DeregisterStreamConsumer",

          //kinesis
          "kinesis:AddTagsToStream",
          "kinesis:CreateStream",
          "kinesis:DeleteStream",
          "kinesis:DescribeStreamSummary",
          "kinesis:EnableEnhancedMonitoring",
          "kinesis:IncreaseStreamRetentionPeriod",
          "kinesis:ListTagsForStream",
          "kinesis:RemoveTagsFromStream",


          //aws_cloudwatch_event_endpoint
          "events:CreateEndpoint",
          "events:DescribeEndpoint",
          "events:CreateEndpoint",
          "events:DeleteEndpoint",
          "events:UpdateEndpoint",
        ],
        "Resource" : "*",
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
