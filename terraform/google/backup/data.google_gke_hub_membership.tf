data "google_gke_hub_membership" "pike" {
  membership_id = "pike"
  location      = "us-central1"
}

output "google_gke_hub_membership" {
  value = data.google_gke_hub_membership.pike
}
