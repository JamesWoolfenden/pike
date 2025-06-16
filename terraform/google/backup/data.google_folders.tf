data "google_folders" "pike" {
  parent_id = "organizations/12345"
}

output "google_folders" {
  value = data.google_folders.pike
}
