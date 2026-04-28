resource "azurerm_healthcare_dicom_service" "pike_gen" {
  name         = "tfexDicom"
  workspace_id = "pike"
  location     = "east us"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    environment = "None"
  }
}
