data "aws_glue_registry" "pike" {
  name = "pike"
}

output "aws_glue_registry" {
  value = data.aws_glue_registry.pike
}
