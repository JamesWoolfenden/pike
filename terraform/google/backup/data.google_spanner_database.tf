data "google_spanner_database" "pike" {
  instance = "pike"
  name     = "pike"
}

output "google_spanner_database" {
  value = data.google_spanner_database.pike
}
