data "google_cloudfunctions2_function" "pike" {
  location = "europe-west2"
  name     = "pike"
}

output "function" {
  value = data.google_cloudfunctions2_function.pike
}
