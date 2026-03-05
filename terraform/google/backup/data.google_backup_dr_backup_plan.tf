data "google_backup_dr_backup_plan" "pike" {
  backup_plan_id = "pike"
  location       = "us-central1"
}

output "google_backup_dr_backup_plan" {
  value = data.google_backup_dr_backup_plan.pike
}
