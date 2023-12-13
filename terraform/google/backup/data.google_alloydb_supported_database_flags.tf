data "google_alloydb_supported_database_flags" "pike" {
  location = "us-central1"
}

output "flags" {
  value = data.google_alloydb_supported_database_flags.pike
}
