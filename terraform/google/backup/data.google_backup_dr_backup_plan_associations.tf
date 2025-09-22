data "google_backup_dr_backup_plan_associations" "pike" {
  resource_type = "sqladmin.googleapis.com/Instance"
  location      = "us-central1"
}

output "google_backup_dr_backup_plan_associations" {
  value = data.google_backup_dr_backup_plan_associations.pike
}
