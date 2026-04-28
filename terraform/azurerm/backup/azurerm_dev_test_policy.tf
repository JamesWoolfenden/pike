resource "azurerm_dev_test_policy" "pike_gen" {
  name                = "LabVmCount"
  policy_set_name     = "default"
  lab_name            = "pike"
  resource_group_name = "pike"
  fact_data           = ""
  threshold           = "999"
  evaluator_type      = "MaxValuePolicy"

  tags = {
    "Acceptance" = "Test"
  }
}
