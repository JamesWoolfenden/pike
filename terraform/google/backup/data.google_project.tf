data "google_project" "pike" {}

output "google_project" {
  value = data.google_project.pike
}
