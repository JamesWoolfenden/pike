data "google_spanner_database_iam_policy" "pike" {
  database = "pike"
  instance = "pike"
}
