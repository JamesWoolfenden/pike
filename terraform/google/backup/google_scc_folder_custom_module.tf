resource "google_scc_folder_custom_module" "pike" {
  folder           = "jim"
  display_name     = "basic_custom_module"
  enablement_state = "ENABLED"
  custom_config {
    predicate {
      expression = "resource.rotationPeriod > duration(\"2592000s\")"
    }
    resource_selector {
      resource_types = [
        "cloudkms.googleapis.com/CryptoKey",
      ]
    }
    description    = "The rotation period of the identified cryptokey resource exceeds 30 days."
    recommendation = "Set the rotation period to at most 30 days."
    severity       = "MEDIUM"
  }
}
