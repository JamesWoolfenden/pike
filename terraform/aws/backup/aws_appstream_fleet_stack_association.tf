resource "aws_appstream_fleet_stack_association" "pike" {
  fleet_name = aws_appstream_fleet.pike.name
  stack_name = aws_appstream_stack.pike.name
}
