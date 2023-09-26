data "google_projects" "pike" {
  filter = "parent.id:012345678910 lifecycleState:DELETE_REQUESTED"
}

output "projects" {
  value = data.google_projects.pike
}
