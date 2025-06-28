data "google_tags_tag_key_iam_policy" "pike" {
}

output "google_tags_tag_key_iam_policy" {
  value = data.google_tags_tag_key_iam_policy.pike
}
