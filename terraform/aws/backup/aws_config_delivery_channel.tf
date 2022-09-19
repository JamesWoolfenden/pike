resource "aws_config_delivery_channel" "pike" {
  s3_bucket_name = "aws-config-config-680235478471"
  //s3_kms_key_arn = ""
  sns_topic_arn = "arn:aws:sns:eu-west-2:680235478471:pike"
}
