resource "azurerm_cognitive_account_rai_policy" "pike_gen" {
  name                 = "example-rai-policy"
  cognitive_account_id = "pike"
  base_policy_name     = "Microsoft.Default"
  content_filter {
    name               = "Hate"
    filter_enabled     = true
    block_enabled      = true
    severity_threshold = "High"
    source             = "Prompt"
  }
}
