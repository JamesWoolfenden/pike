data "google_storage_bucket_objects" "pike" {
  bucket = "stuffidontwanttoshare"
}

output "google_storage_bucket_objects" {
  value = data.google_storage_bucket_objects.pike
}
