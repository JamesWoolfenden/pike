resource "google_secret_manager_regional_secret" "pike" {
  secret_id = "tf-reg-secret"
  location  = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}
