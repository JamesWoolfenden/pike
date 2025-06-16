data "google_dataplex_aspect_type_iam_policy" "pike" {
  aspect_type_id = "pike"
}

output "google_dataplex_aspect_type_iam_policy" {
  value = data.google_dataplex_aspect_type_iam_policy.pike
}
