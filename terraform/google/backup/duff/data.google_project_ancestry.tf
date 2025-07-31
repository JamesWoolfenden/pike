data "google_project_ancestry" "pike" {
}

output "google_project_ancestry" {
  value = data.google_project_ancestry.pike
}
