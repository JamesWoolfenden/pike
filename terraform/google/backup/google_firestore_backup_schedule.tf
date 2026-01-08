resource "google_firestore_backup_schedule" "pike" {
  database  = google_firestore_database.pike.name
  retention = "8467200s" // 14 weeks (maximum possible retention)
  daily_recurrence {}
}
