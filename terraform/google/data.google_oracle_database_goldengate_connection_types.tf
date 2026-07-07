data "google_oracle_database_goldengate_connection_types" "pike" {
  location = "us-east4"
}

output "google_oracle_database_goldengate_connection_types" {
  value = data.google_oracle_database_goldengate_connection_types.pike
}
