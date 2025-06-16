data "google_storage_buckets" "pike" {
}

output "google_storage_buckets" {
  value = data.google_storage_buckets.pike
}
