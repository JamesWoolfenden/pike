resource "aws_connect_queue" "pike" {
  instance_id = aws_connect_instance.pike.id
  name        = "Example Name"

  description           = "Example Description"
  hours_of_operation_id = aws_connect_hours_of_operation.pike.hours_of_operation_id

  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
