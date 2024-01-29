resource "google_alloydb_backup" "default" {
  location     = "europe-west2"
  backup_id    = "alloydb-backup"
  cluster_name = google_alloydb_cluster.default.name

  depends_on = [google_alloydb_instance.default]
}
