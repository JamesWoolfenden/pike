resource "google_bigtable_logical_view" "logical_view" {
  logical_view_id     = "btlogical"
  instance            = google_bigtable_instance.instance.name
  deletion_protection = false
  query               = <<EOT
SELECT _key,
FROM tf-table;
EOT

  depends_on = [
    google_bigtable_table.table
  ]
}
