data "aws_servicecatalogappregistry_attribute_group_associations" "pike" {
  name = "pike"
}

output "aws_servicecatalogappregistry_attribute_group_associations" {
  value = data.aws_servicecatalogappregistry_attribute_group_associations.pike
}
