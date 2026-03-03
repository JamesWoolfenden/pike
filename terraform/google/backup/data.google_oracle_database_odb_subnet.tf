data "google_oracle_database_odb_subnet" "pike" {
}

output "google_oracle_database_odb_subnet" {
  value = data.google_oracle_database_odb_subnet.pike
}
