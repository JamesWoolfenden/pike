data "google_sql_database_instances" "pike" {}

output "instances" {
  value = data.google_sql_database_instances.pike
}
