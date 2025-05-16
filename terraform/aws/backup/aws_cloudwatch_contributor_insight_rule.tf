resource "aws_cloudwatch_contributor_insight_rule" "pike" {
  rule_name = "pike"
  rule_definition = jsonencode({
    "Schema" : {
      "Name" : "CloudWatchLogRule",
      "Version" : 1
    },
    "LogGroupNames" : [
      "/aws/containerinsights/sample-cluster-name/flowlogs"
    ],
    "LogFormat" : "CLF",
    "Fields" : {
      "4" : "srcaddr",
      "5" : "dstaddr",
      "10" : "bytes"
    },
    "Contribution" : {
      "Keys" : [
        "srcaddr",
        "dstaddr"
      ],
      "ValueOf" : "bytes",
      "Filters" : []
    },
    "AggregateOn" : "Sum"
    }

  )
  tags = {
    pike = "permission"
    # delete = "me"
  }
}
