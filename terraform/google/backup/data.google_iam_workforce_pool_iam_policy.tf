data "google_iam_workforce_pool_iam_policy" "pike" {
}

output "google_iam_workforce_pool_iam_policy" {
  value = data.google_iam_workforce_pool_iam_policy.pike
}
