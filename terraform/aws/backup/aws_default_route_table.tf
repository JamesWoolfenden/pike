resource "aws_default_route_table" "pike" {
  default_route_table_id = "rtb-000df9b1b04b02616"
  tags = {
    Name = "Default Route Table"
    pike = "permissions"
  }
}

output "default_route_table" {
  value = aws_default_route_table.pike
}
