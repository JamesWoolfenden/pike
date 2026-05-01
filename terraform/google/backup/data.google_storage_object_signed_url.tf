data "google_storage_object_signed_url" "pike" {
  bucket = "pike-example"
  path   = "./"
}

output "google_storage_object_signed_url" {
  value     = data.google_storage_object_signed_url.pike
  sensitive = true
}
