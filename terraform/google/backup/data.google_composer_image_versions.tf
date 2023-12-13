data "google_composer_image_versions" "pike" {}

output "versions" {
  value = data.google_composer_image_versions.pike
}
