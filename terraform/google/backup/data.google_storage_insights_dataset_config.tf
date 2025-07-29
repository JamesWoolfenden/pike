data "google_storage_insights_dataset_config" "pike" {
}

output "google_storage_insights_dataset_config" {
  value = data.google_storage_insights_dataset_config.pike
}
