resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "bigquery.connections.create",
    "bigquery.connections.delete",
    "bigquery.connections.get",
    "bigquery.connections.update",
    "storage.buckets.get"
  ]
}
