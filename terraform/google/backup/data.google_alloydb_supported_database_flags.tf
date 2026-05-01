data "google_alloydb_supported_database_flags" "pike" {
  location = "us-central1"
}

output "google_alloydb_supported_database_flags" {
  value = data.google_alloydb_supported_database_flags.pike
}
