data "aws_servicecatalogappregistry_application" "pike" {
  id = "application-1234"
}

output "aws_servicecatalogappregistry_application" {
  value = data.aws_servicecatalogappregistry_application.pike
}
