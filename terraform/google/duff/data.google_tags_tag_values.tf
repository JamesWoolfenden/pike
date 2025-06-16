data "google_tags_tag_values" "pike" {
  parent = "pike"
}

output "google_tags_tag_values" {
  value = data.google_tags_tag_values.pike
}
