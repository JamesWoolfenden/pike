resource "google_biglake_iceberg_catalog" "pike" {
  name         = "my_iceberg_catalog"
  catalog_type = "CATALOG_TYPE_GCS_BUCKET"
}
