data "google_dataplex_data_quality_rules" "pike" {
  project      = "pike-412922"
  location     = "us-central1"
  data_scan_id = "pike"
}

output "google_dataplex_data_quality_rules" {
  value = data.google_dataplex_data_quality_rules.pike
}
