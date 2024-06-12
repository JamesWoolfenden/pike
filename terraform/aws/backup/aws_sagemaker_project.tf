resource "aws_sagemaker_project" "pike" {
  project_name = "pike"

  service_catalog_provisioning_details {
    product_id = aws_servicecatalog_product.example.id
  }

  tags = {
    pike = "permissions"
  }
}

resource "aws_servicecatalog_product" "example" {
  name  = "example"
  owner = "example-owner"
  type  = "CLOUD_FORMATION_TEMPLATE"

  provisioning_artifact_parameters {
    template_url = "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json"
  }

  tags = {
    foo = "bar"
  }
}
