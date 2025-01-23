resource "aws_appstream_user_stack_association" "pike" {
  authentication_type = aws_appstream_user.pike.authentication_type
  stack_name          = aws_appstream_stack.pike.name
  user_name           = aws_appstream_user.pike.user_name
}
