data "google_gemini_repository_group_iam_policy" "pike" {
  code_repository_index = "pike"
  repository_group_id   = "pike"
  project               = "pike-412922"
}

output "google_gemini_repository_group_iam_policy" {
  value = data.google_gemini_repository_group_iam_policy.pike
}
