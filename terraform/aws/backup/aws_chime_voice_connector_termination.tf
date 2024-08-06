resource "aws_chime_voice_connector_termination" "pike" {
  disabled           = false
  cps_limit          = 1
  cidr_allow_list    = ["50.35.78.96/31"]
  calling_regions    = ["US", "CA"]
  voice_connector_id = aws_chime_voice_connector.pike.id
}
