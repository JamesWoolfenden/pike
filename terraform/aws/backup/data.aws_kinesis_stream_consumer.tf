data "aws_kinesis_stream_consumer" "pike" {
  stream_arn = "arn:aws:kinesis:eu-west-2:${data.aws_caller_identity.current.account_id}:stream/tablename"
}

data "aws_caller_identity" "current" {}
