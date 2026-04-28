resource "azurerm_machine_learning_inference_cluster" "pike_gen" {
  name                  = "example"
  location              = "pike"
  cluster_purpose       = "FastProd"
  kubernetes_cluster_id = "pike"
  description           = "This is an example cluster used with Terraform"

  machine_learning_workspace_id = "pike"

  tags = {
    "stage" = "example"
  }
}
