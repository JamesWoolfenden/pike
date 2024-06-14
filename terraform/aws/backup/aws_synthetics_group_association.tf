resource "aws_synthetics_group_association" "pike" {
  group_name = "pike"
  canary_arn = aws_synthetics_canary.pike.arn
}
