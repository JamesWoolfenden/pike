data "google_folder" "pike" {
  folder = "pike"
}

output "google_folder" {
  value = data.google_folder.pike
}
