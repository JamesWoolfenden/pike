data "google_sql_backup_run" "pike" {
  instance = "examplea"
}

output "run" {
  value = data.google_sql_backup_run.pike
}
