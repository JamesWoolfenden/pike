resource "aws_connect_instance_storage_config" "pike" {
  instance_id   = aws_connect_instance.pike.id
  resource_type = "CONTACT_TRACE_RECORDS"

  storage_config {
    kinesis_firehose_config {
      firehose_arn = aws_kinesis_firehose_delivery_stream.pike.arn
    }
    storage_type = "KINESIS_FIREHOSE"
  }

}


resource "aws_kinesis_firehose_delivery_stream" "pike" {
  name        = "terraform-kinesis-firehose-extended-s3-pike-stream2"
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
