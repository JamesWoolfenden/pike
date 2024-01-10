resource "aws_servicecatalog_provisioned_product" "pike" {
  name                       = "example"
  product_name               = "Example product"
  provisioning_artifact_name = "Example version"

  provisioning_parameters {
    key   = "foo"
    value = "bar"
  }

  tags = {
    foo = "bar"
  }
}
