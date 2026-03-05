data "google_cloud_asset_search_all_resources" "pike" {
  scope = "projects/pike"
}

output "google_cloud_asset_search_all_resources" {
  value = data.google_cloud_asset_search_all_resources.pike
}
