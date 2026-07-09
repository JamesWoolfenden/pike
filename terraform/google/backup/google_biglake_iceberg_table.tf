resource "google_biglake_iceberg_table" "pike" {
  name      = "pike"
  catalog   = google_biglake_iceberg_catalog.pike.name
  namespace = google_biglake_iceberg_namespace.pike.namespace_id

  schema {
    fields {
      id       = 1
      name     = "id"
      type     = "long"
      required = true
    }
  }
  depends_on = [
    google_biglake_iceberg_namespace.pike
  ]
}
