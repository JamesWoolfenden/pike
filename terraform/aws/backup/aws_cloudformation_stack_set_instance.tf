resource "aws_cloudformation_stack_set_instance" "pike" {
  account_id     = "123456789012"
  region         = "eu-west-2"
  stack_set_name = aws_cloudformation_stack_set.example.name
}