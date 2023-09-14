resource "google_sql_database" "pike" {
  name     = "pike"
  instance = google_sql_database_instance.main.name
}
