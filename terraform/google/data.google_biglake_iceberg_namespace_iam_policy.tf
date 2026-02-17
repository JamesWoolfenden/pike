data "google_biglake_iceberg_namespace_iam_policy" "pike" {
  //provider = google-beta
}

output "google_biglake_iceberg_namespace_iam_policy" {
  value = data.google_biglake_iceberg_namespace_iam_policy.pike
}
