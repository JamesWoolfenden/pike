resource "google_data_catalog_policy_tag" "pike" {
  taxonomy     = google_data_catalog_taxonomy.pike.id
  display_name = "Low security"
  description  = "A policy tag normally associated with low security items"
}
