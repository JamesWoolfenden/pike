resource "aws_imagebuilder_infrastructure_configuration" "pike" {
  description                   = "pike description"
  instance_profile_name         = "imagebuilder_20220921141023768700000002"
  instance_types                = ["t2.nano", "t3.micro"]
  key_pair                      = "test"
  name                          = "pike"
  security_group_ids            = ["sg-06db45d8099f7f549"]
  sns_topic_arn                 = "arn:aws:sns:eu-west-2:680235478471:pike"
  subnet_id                     = "subnet-08d97e381dbc80d40"
  terminate_instance_on_failure = true

  logging {
    s3_logs {
      s3_bucket_name = "logging-680235478471"
      s3_key_prefix  = "logs"
    }
  }


  tags = {
    pike = "permissions"
    foo  = "bar"
  }
}
