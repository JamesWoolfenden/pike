data "aws_servicecatalogappregistry_attribute_group" "pike" {
  name = "pike"
}

output "aws_servicecatalogappregistry_attribute_group" {
  value = data.aws_servicecatalogappregistry_attribute_group.pike
}
