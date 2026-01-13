resource "google_assured_workloads_workload" "pike" {
  compliance_regime = "FEDRAMP_MODERATE"
  display_name      = "{{display}}"
  location          = "us-west1"
  organization      = "123456789"
  billing_account   = "billingAccounts/000000-0000000-0000000-000000"

  kms_settings {
    next_rotation_time = "9999-10-02T15:01:23Z"
    rotation_period    = "10368000s"
  }

  provisioned_resources_parent = "folders/519620126891"

  resource_settings {
    display_name  = "{{name}}"
    resource_type = "CONSUMER_FOLDER"
  }

  resource_settings {
    resource_type = "ENCRYPTION_KEYS_PROJECT"
  }

  resource_settings {
    resource_id   = "ring"
    resource_type = "KEYRING"
  }

  violation_notifications_enabled = true

  workload_options {
    kaj_enrollment_type = "KEY_ACCESS_TRANSPARENCY_OFF"
  }

  labels = {
    label-one = "value-one"
  }
}
