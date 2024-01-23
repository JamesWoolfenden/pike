data "google_sql_database_instance_latest_recovery_time" "pike" {
  instance = "examplea"
}

output "time" {
  value = data.google_sql_database_instance_latest_recovery_time.pike
}
