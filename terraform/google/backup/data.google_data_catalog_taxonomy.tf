data "google_data_catalog_taxonomy" "pike" {
  display_name = "pike-test-taxonomy"
  region       = "europe-west2"
}

output "google_data_catalog_taxonomy" {
  value = data.google_data_catalog_taxonomy.pike
}