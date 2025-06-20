resource "google_spanner_instance_config" "example" {
  name         = "custom-nam11-config"
  display_name = "Test Spanner Instance Config"
  base_config  = "nam11"
  replicas {
    location                = "us-west1"
    type                    = "READ_ONLY"
    default_leader_location = false
  }
  labels = {
    "foo" = "bar"
  }
}
