data "google_oracle_database_autonomous_databases" "pike" {
  location = "us-central1"
}

output "google_oracle_database_autonomous_databases" {
  value = data.google_oracle_database_autonomous_databases.pike
}
