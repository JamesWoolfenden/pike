data "google_tags_tag_keys" "pike" {
  parent = "pike"
}

output "google_tags_tag_keys" {
  value = data.google_tags_tag_keys.pike
}
