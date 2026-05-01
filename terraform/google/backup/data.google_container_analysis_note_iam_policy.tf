data "google_container_analysis_note_iam_policy" "pike" {
  note = "pike"
}

output "google_container_analysis_note_iam_policy" {
  value = data.google_container_analysis_note_iam_policy.pike
}
