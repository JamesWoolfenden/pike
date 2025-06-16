data "google_dataplex_entry_type_iam_policy" "pike" {
  entry_type_id = "pike"
}

output "google_dataplex_entry_type_iam_policy" {
  value = data.google_dataplex_entry_type_iam_policy.pike
}
