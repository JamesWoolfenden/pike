data "aws_default_tags" "pike" {
}

output "tags" {
  value = data.aws_default_tags.pike
}
