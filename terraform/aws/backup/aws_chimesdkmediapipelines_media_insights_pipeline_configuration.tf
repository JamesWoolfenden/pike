resource "aws_chimesdkmediapipelines_media_insights_pipeline_configuration" "pike" {
  name                     = "MyBasicConfiguration"
  resource_access_role_arn = aws_iam_role.call_analytics_role.arn
  elements {
    type = "AmazonTranscribeCallAnalyticsProcessor"
    amazon_transcribe_call_analytics_processor_configuration {
      language_code = "en-US"
    }
  }

  elements {
    type = "KinesisDataStreamSink"
    kinesis_data_stream_sink_configuration {
      insights_target = aws_kinesis_stream.example.arn
    }
  }

  tags = {
    Key1 = "Value1"
    Key2 = "Value2"
  }
}

resource "aws_iam_role" "call_analytics_role" {
  name = "call_analytics_role"
  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : "sts:AssumeRole",
        "Principal" : {
          "Service" : "chime.amazonaws.com"
        },
        "Effect" : "Allow",
        "Sid" : ""
      }
    ]
  })
}

resource "aws_kinesis_stream" "example" {
  shard_count      = 1
  retention_period = 48

  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]

  stream_mode_details {
    stream_mode = "PROVISIONED"
  }

  tags = {
    Environment = "test"
  }
  name = "pike"
}
