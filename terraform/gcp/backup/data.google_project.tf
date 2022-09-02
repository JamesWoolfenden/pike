data "google_project" "pike" {}

output "project" {
  value = data.google_project.pike
}
