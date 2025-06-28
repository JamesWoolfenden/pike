data "google_organization_iam_policy" "pike" {
}

output "google_organization_iam_policy" {
  value = data.google_organization_iam_policy.pike
}
