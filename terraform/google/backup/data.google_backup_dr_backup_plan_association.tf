data "google_backup_dr_backup_plan_association" "pike" {
  location                   = "us-central1"
  backup_plan_association_id = "pike"
}

output "google_backup_dr_backup_plan_association" {
  value = data.google_backup_dr_backup_plan_association.pike
}
