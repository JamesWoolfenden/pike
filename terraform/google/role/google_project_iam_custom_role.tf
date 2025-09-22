
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "artifactregistry.pythonpackages.list",
    "bigquery.dataPolicies.getIamPolicy",
    "iap.webServices.getIamPolicy",
    "backupdr.dataSourceReferences.fetchForAlloydbCluster",
    "backupdr.dataSourceReferences.fetchForCloudSqlInstance",
    "backupdr.backupPlanAssociations.fetchForAlloydbCluster",
    "backupdr.backupPlanAssociations.fetchForCloudSqlInstance",
    "backupdr.backupPlanAssociations.fetchForComputeDisk",
    "backupdr.backupPlanAssociations.fetchForComputeInstance",
    "backupdr.backupPlanAssociations.getForAlloydbCluster",
    "backupdr.backupPlanAssociations.getForCloudSqlInstance",
    "backupdr.backupPlanAssociations.getForComputeDisk",
    "backupdr.backupPlanAssociations.getForComputeInstance"
  ]
}
