data "google_oracle_database_autonomous_database" "pike" {
  location               = "us-central1"
  autonomous_database_id = "pike"
}

output "google_oracle_database_autonomous_database" {
  value = data.google_oracle_database_autonomous_database.pike
}
