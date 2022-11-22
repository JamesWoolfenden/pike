resource "random_string" "pike" {
  length           = 16
  special          = true
  override_special = "/@£$"
}
