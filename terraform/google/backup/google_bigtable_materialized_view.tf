resource "google_bigtable_materialized_view" "materialized_view" {
  materialized_view_id = "btmaterialized"
  instance             = google_bigtable_instance.instance.name
  deletion_protection  = false
  query                = <<EOT
SELECT _key,
FROM tf-table
ORDER BY _key;
EOT

  depends_on = [
    google_bigtable_table.table
  ]
}
