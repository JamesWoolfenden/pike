data "aws_default_tags" "pike" {
}

output "default_tags" {
  value = data.aws_default_tags.pike
}
