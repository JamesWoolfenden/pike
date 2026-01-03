data "google_gke_hub_membership_binding" "pike" {
  membership_id         = "pike"
  membership_binding_id = "pike"
  location              = "global"
}

output "google_gke_hub_membership_binding" {
  value = data.google_gke_hub_membership_binding.pike
}
