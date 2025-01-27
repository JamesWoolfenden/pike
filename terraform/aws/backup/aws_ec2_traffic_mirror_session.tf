resource "aws_ec2_traffic_mirror_session" "pike" {
  description              = "traffic mirror session - terraform example"
  network_interface_id     = aws_instance.test.primary_network_interface_id
  session_number           = 1
  traffic_mirror_filter_id = aws_ec2_traffic_mirror_filter.filter.id
  traffic_mirror_target_id = aws_ec2_traffic_mirror_target.pike.id
}

resource "aws_ec2_traffic_mirror_filter" "filter" {
  description      = "traffic mirror filter - terraform example"
  network_services = ["amazon-dns"]
}
