resource "azurerm_security_center_contact" "pike" {
  name  = "James"
  email = "jim_woolfenden@hotmail.com"
  phone = "07976379686"

  alert_notifications = true
  alerts_to_admins    = true
}
