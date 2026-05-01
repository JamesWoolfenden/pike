data "google_sourcerepo_repository_iam_policy" "pike" {
  repository = "pike"
}

output "google_sourcerepo_repository_iam_policy" {
  value = data.google_sourcerepo_repository_iam_policy.pike
}
