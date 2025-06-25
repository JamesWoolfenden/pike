resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "analyticshub.dataExchanges.create",
    "analyticshub.dataExchanges.delete",
    "analyticshub.dataExchanges.get",
    "analyticshub.dataExchanges.update",
    "analyticshub.listings.create",
    "analyticshub.listings.delete",
    "analyticshub.listings.get",
    "analyticshub.listings.update",
    "bigquery.datasets.create",
    "bigquery.datasets.delete",
    "bigquery.datasets.get",
    "bigquery.datasets.update",
    "bigtable.clusters.list",
    "bigtable.instances.create",
    "bigtable.instances.delete",
    "bigtable.instances.get",
    "bigtable.instances.list",
    "bigtable.instances.update",
    "bigtable.tables.create",
    "bigtable.tables.delete",
    "bigtable.tables.get",
    "iam.serviceAccounts.create",
    "iam.serviceAccounts.delete",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.update",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.update",
    "storage.objects.create",
    "storage.objects.delete",
    "storage.objects.get",

    //google_bigquery_analytics_hub_listing_subscription

    //google_bigtable_instance
    "bigtable.clusters.create",
    "bigtable.clusters.delete",
    "bigtable.clusters.get",
    "bigtable.clusters.update",



    //google_biglake_database


    //google_biglake_table

  ]
}
