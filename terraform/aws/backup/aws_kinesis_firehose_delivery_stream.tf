resource "aws_kinesis_firehose_delivery_stream" "extended_s3_stream" {
  name        = "terraform-kinesis-firehose-extended-s3-pike-stream"
  destination = "extended_s3"

  extended_s3_configuration {
    role_arn   = "arn:aws:iam::680235478471:role/firehose"
    bucket_arn = "arn:aws:s3:::testbucketforlbjgw"

    processing_configuration {
      enabled = "true"

      processors {
        type = "Lambda"

        parameters {
          parameter_name  = "LambdaArn"
          parameter_value = "arn:aws:lambda:eu-west-2:680235478471:function:updatepolicy:$LATEST"
        }
      }
    }
  }
  tags = {
    pike = "permissions"
  }
}
