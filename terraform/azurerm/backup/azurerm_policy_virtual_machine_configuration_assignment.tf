resource "azurerm_policy_virtual_machine_configuration_assignment" "pike_gen" {
  name               = "AzureWindowsBaseline"
  location           = "pike"
  virtual_machine_id = "pike"

  configuration {
    assignment_type = "ApplyAndMonitor"
    version         = "1.*"

    parameter {
      name  = "Minimum Password Length;ExpectedValue"
      value = "16"
    }
    parameter {
      name  = "Minimum Password Age;ExpectedValue"
      value = "0"
    }
    parameter {
      name  = "Maximum Password Age;ExpectedValue"
      value = "30,45"
    }
    parameter {
      name  = "Enforce Password History;ExpectedValue"
      value = "10"
    }
    parameter {
      name  = "Password Must Meet Complexity Requirements;ExpectedValue"
      value = "1"
    }
  }
}
