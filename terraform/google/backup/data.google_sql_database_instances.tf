data "google_sql_database_instances" "pike" {}

output "google_sql_database_instances" {
  value = data.google_sql_database_instances.pike
}
