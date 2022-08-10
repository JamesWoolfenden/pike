resource "aws_cloudwatch_log_subscription_filter" "test_lambdafunction_logfilter" {
  name = "pike_lambdafunction_logfilter"
  //role_arn        = "arn:aws:iam::680235478471:role/pike-test-role"
  log_group_name  = "/aws/lambda/reads3"
  filter_pattern  = "logtype test2"
  destination_arn = "arn:aws:lambda:eu-west-2:680235478471:function:updatepolicy"
  distribution    = "Random"
}
