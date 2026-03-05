data "google_storage_managed_folder_iam_policy" "pike" {
  bucket         = google_storage_bucket.pike_managed_folder_test.name
  managed_folder = google_storage_managed_folder.pike.name
}

output "google_storage_managed_folder_iam_policy" {
  value = data.google_storage_managed_folder_iam_policy.pike
}
