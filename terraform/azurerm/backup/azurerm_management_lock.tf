resource "azurerm_management_lock" "pike_gen" {
  name       = "subscription-level"
  scope      = "pike"
  lock_level = "CanNotDelete"
  notes      = "Items can't be deleted in this subscription!"
}
