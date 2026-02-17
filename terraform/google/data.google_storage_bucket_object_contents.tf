data "google_storage_bucket_object_contents" "pike" {
  //provider = google-beta
}

output "google_storage_bucket_object_contents" {
  value = data.google_storage_bucket_object_contents.pike
}
