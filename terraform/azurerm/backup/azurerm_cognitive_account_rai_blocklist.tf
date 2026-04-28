resource "azurerm_cognitive_account_rai_blocklist" "pike_gen" {
  name                 = "example-crb"
  cognitive_account_id = "pike"
  description          = "Azure OpenAI Rai Blocklist"
}
