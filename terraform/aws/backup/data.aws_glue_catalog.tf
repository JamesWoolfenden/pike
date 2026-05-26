data "aws_glue_catalog" "pike" {
  name = "pike"
}

output "aws_glue_catalog" {
  value = data.aws_glue_catalog.pike
}
