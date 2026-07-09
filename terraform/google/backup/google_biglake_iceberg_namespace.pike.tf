resource "google_biglake_iceberg_namespace" "pike" {
  catalog      = google_biglake_iceberg_catalog.pike.name
  namespace_id = "pike"
}
