data "google_dataplex_entry_group_iam_policy" "pike" {
  entry_group_id = "pike"
}

output "google_dataplex_entry_group_iam_policy" {
  value = data.google_dataplex_entry_group_iam_policy.pike
}
