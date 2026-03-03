data "google_oracle_database_odb_network" "pike" {
}

output "google_oracle_database_odb_network" {
  value = data.google_oracle_database_odb_network.pike
}
