provider "azurerm" {
  features {
    cognitive_account {
      purge_soft_delete_on_destroy = true
    }

    key_vault {
      purge_soft_delete_on_destroy = true
    }
  }
  subscription_id = "037ce662-dfc1-4b8b-a8a7-6c414b540ed6"
}
