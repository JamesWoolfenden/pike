data "google_sql_backup_run" "pike" {
  instance = "examplea"
}

output "google_sql_backup_run" {
  value = data.google_sql_backup_run.pike
}
