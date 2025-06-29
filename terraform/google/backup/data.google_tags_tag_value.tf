data "google_tags_tag_value" "pike" {
}

output "google_tags_tag_value" {
  value = data.google_tags_tag_value.pike
}
