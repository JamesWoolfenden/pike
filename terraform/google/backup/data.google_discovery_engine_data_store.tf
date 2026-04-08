data "google_discovery_engine_data_store" "pike" {
  display_name = "pike"
}

output "google_discovery_engine_data_store" {
  value = data.google_discovery_engine_data_store.pike
}
