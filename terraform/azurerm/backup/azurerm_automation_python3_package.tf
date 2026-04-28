resource "azurerm_automation_python3_package" "pike_gen" {
  name                    = "example"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  content_uri             = "https://pypi.org/packages/source/r/requests/requests-2.31.0.tar.gz"
  content_version         = "2.31.0"
  hash_algorithm          = "sha256"
  hash_value              = "942c5a758f98d790eaed1a29cb6eefc7ffb0d1cf7af05c3d2791656dbd6ad1e1"
  tags = {
    key = "foo"
  }
}
