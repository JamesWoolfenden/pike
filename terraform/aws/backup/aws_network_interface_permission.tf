resource "aws_network_interface_permission" "pike" {
  network_interface_id = aws_network_interface.pike.id
  aws_account_id       = "123456789012"
  permission           = "INSTANCE-ATTACH"
}


resource "aws_network_interface" "pike" {
  subnet_id = "subnet-09ff91b5b0adb1fd4"
}
