data "google_gke_hub_feature" "pike" {
  name     = "servicemesh"
  location = "global"
}

output "google_gke_hub_feature" {
  value = data.google_gke_hub_feature.pike
}
