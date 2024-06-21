resource "aws_oam_link" "pike" {
  label_template  = "$AccountName"
  resource_types  = ["AWS::CloudWatch::Metric"]
  sink_identifier = aws_oam_sink.pike.id
  tags = {
    Env  = "prod"
    pike = "permissions"
  }
}
