resource "aws_ec2_network_insights_path" "pike" {
  source      = aws_network_interface.source.id
  destination = aws_network_interface.destination.id
  protocol    = "tcp"
  tags = {
    pike = "permissions"
  }
}

resource "aws_ec2_network_insights_analysis" "analysis" {
  network_insights_path_id = aws_ec2_network_insights_path.pike.id
  tags = {
    pike = "permissions"
  }
}

resource "aws_network_interface" "source" {
  subnet_id = "subnet-09ff91b5b0adb1fd4"
}

resource "aws_network_interface" "destination" {
  subnet_id = "subnet-09ff91b5b0adb1fd4"
}
