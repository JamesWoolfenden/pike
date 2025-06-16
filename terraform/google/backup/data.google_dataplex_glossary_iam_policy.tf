data "google_dataplex_glossary_iam_policy" "pike" {
  glossary_id = "pike"
}

output "google_dataplex_glossary_iam_policy" {
  value = data.google_dataplex_glossary_iam_policy.pike
}
