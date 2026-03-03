data "google_biglake_iceberg_namespace_iam_policy" "pike" {
}

output "google_biglake_iceberg_namespace_iam_policy" {
  value = data.google_biglake_iceberg_namespace_iam_policy.pike
}
