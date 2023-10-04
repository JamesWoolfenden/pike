data "google_bigtable_table_iam_policy" "pike" {
  table    = "pike"
  instance = "pike"
}
