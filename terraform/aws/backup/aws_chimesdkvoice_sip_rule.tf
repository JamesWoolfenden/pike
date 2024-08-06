resource "aws_chimesdkvoice_sip_rule" "pike" {
  provider      = aws.central
  name          = "example-sip-rule"
  trigger_type  = "RequestUriHostname"
  trigger_value = aws_chime_voice_connector.pike.outbound_host_name
  target_applications {
    priority                 = 1
    sip_media_application_id = aws_chimesdkvoice_sip_media_application.pike.id
    aws_region               = "us-east-1"
  }

}
