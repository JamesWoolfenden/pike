resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "bigquery.datasets.create",
    "cloudscheduler.jobs.create",
    "cloudscheduler.jobs.delete",
    "cloudscheduler.jobs.enable",
    "cloudscheduler.jobs.get",
    "cloudscheduler.jobs.update",
    "iam.serviceAccounts.create",
    "iam.serviceAccounts.delete",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.update",
    "resourcemanager.projects.get",
    "resourcemanager.projects.getIamPolicy",
    "resourcemanager.projects.setIamPolicy",
    "run.jobs.create",
    "run.jobs.delete",
    "run.jobs.get",
    "run.jobs.update",
    "run.operations.get",

    //google_bigquery_table_iam_policy,google_bigquery_table_iam_member,google_bigquery_table_iam_binding
    "bigquery.tables.getIamPolicy",
    "bigquery.tables.setIamPolicy",

    //google_bigquery_reservation
    "bigquery.reservations.create",
    "bigquery.reservations.update",
    "bigquery.reservations.delete",
    "bigquery.reservations.get",

    //google_bigquery_connection
    "bigquery.connections.create",
    "bigquery.connections.update",
    "bigquery.connections.delete",
    "bigquery.connections.get",

    //google_bigquery_capacity_commitment
    "bigquery.capacityCommitments.create",
    "bigquery.capacityCommitments.delete",
    "bigquery.capacityCommitments.update",
    "bigquery.capacityCommitments.get",

    //google_bigquery_bi_reservation
    "bigquery.bireservations.update",
    "bigquery.bireservations.get",

    //google_bigquery_analytics_hub_data_exchange
    "analyticshub.dataExchanges.create",
    "analyticshub.dataExchanges.delete",
    "analyticshub.dataExchanges.update",
    "analyticshub.dataExchanges.get",

    //google_bigquery_reservation_assignment
    "bigquery.reservationAssignments.create",
    "bigquery.reservationAssignments.delete",
    "bigquery.reservationAssignments.update",
    "bigquery.reservationAssignments.get",
    "bigquery.reservationAssignments.list",

    //google_bigquery_dataset
    "bigquery.datasets.get",
    "bigquery.datasets.update",

    //google_bigquery_routine
    "bigquery.routines.get",

    //google_bigquery_table
    "bigquery.tables.get",

    //google_bigquery_dataset_iam_binding,google_bigquery_dataset_iam_member
    "bigquery.datasets.update"

  ]
}
