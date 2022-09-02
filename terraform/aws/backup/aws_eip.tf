resource "aws_eip" "example" {
  //instance = "i-0b17edfc889671ee1"
  network_interface = "eni-0edf077fa3bf6a7db"
  tags = {
    pike = "permissions"
  }
}
