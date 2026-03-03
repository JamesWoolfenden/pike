data "google_storage_bucket_object_contents" "pike" {
  bucket = "terraform-pike-bucket-tfstate"
}

output "google_storage_bucket_object_contents" {
  value = data.google_storage_bucket_object_contents.pike
}
