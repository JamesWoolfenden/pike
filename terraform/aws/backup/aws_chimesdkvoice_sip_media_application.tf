resource "aws_chimesdkvoice_sip_media_application" "pike" {
  provider   = aws.central
  aws_region = "eu-west-2"
  name       = "example-sip-media-application"
  endpoints {
    lambda_arn = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
  }
  tags = {
    pike = "permissions"
  }
}
