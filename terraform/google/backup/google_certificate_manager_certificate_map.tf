resource "google_certificate_manager_certificate_map" "default" {
  name        = "cert-map"
  description = "My acceptance test certificate map"
  labels = {
    "terraform" : true,
    "acc-test" : true,
  }
}
