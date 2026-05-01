data "google_composer_image_versions" "pike" {}

output "google_composer_image_versions" {
  value = data.google_composer_image_versions.pike
}
