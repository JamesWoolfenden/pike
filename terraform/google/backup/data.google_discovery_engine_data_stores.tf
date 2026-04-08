data "google_discovery_engine_data_stores" "pike" {
}

output "google_discovery_engine_data_stores" {
  value = data.google_discovery_engine_data_stores.pike
}
