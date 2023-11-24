data "google_container_analysis_note_iam_policy" "pike" {
  note = "pike"
}

output "policy" {
  value = data.google_container_analysis_note_iam_policy.pike
}
