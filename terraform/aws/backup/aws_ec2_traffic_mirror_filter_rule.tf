resource "aws_ec2_traffic_mirror_filter_rule" "pike" {
  description              = "test rule"
  traffic_mirror_filter_id = aws_ec2_traffic_mirror_filter.pike.id
  destination_cidr_block   = "10.0.0.0/8"
  source_cidr_block        = "10.0.0.0/8"
  rule_number              = 1
  rule_action              = "accept"
  traffic_direction        = "egress"
}
