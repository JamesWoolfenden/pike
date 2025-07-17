
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    // google_backup_dr_backup_plan
    "backupdr.backupPlans.create",
    "backupdr.backupPlans.delete",
    "backupdr.backupPlans.get",
    "backupdr.backupPlans.update",

    // google_backup_dr_backup_plan_association
    "backupdr.backupPlanAssociations.createForCloudSqlInstance",
    "backupdr.backupPlanAssociations.createForComputeDisk",
    "backupdr.backupPlanAssociations.createForComputeInstance",
    "backupdr.backupPlanAssociations.deleteForCloudSqlInstance",
    "backupdr.backupPlanAssociations.deleteForComputeDisk",
    "backupdr.backupPlanAssociations.deleteForComputeInstance",
    "backupdr.backupPlanAssociations.fetchForCloudSqlInstance",
    "backupdr.backupPlanAssociations.get",
    "backupdr.backupPlanAssociations.getForCloudSqlInstance",
    "backupdr.backupPlanAssociations.getForComputeDisk",
    "backupdr.backupPlanAssociations.triggerBackupForCloudSqlInstance",
    "backupdr.backupPlanAssociations.triggerBackupForComputeDisk",
    "backupdr.backupPlanAssociations.triggerBackupForComputeInstance",
    "backupdr.backupPlanAssociations.updateForComputeDisk",
    "backupdr.backupPlanAssociations.updateForComputeInstance",

    // google_backup_dr_backup_vault
    "backupdr.backupVaults.create",
    "backupdr.backupVaults.delete",
    "backupdr.backupVaults.get",
    "backupdr.backupVaults.update",
    // google_backup_dr_management_server
    "backupdr.managementServers.create",
    "backupdr.managementServers.get",
    "backupdr.managementServers.update",
    "backupdr.managementServers.delete",
    // google_backup_dr_service_config
    "backupdr.resourceBackupConfigs.get",
  ]
}
