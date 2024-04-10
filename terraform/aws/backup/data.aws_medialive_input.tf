data "aws_medialive_input" "pike" {
  id = "12345"
}

output "aws_medialive_input" {
  value = data.aws_medialive_input.pike
}
