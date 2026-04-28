resource "azurerm_api_management_api_tag_description" "pike_gen" {
  api_tag_id                = "pike"
  description               = "This is an example description"
  external_docs_url         = "https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs"
  external_docs_description = "This is an example external docs description"
}
