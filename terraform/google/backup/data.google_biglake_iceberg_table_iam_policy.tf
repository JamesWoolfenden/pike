data "google_biglake_iceberg_table_iam_policy" "pike" {
  name      = "pike"
  catalog   = "pike"
  namespace = "pike"
}

output "google_biglake_iceberg_table_iam_policy" {
  value = data.google_biglake_iceberg_table_iam_policy.pike
}
