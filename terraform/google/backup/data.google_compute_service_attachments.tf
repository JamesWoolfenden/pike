data "google_compute_service_attachments" "pike" {
}

output "google_compute_service_attachments" {
  value = data.google_compute_service_attachments.pike
}