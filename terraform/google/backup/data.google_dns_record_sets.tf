data "google_dns_record_sets" "pike" {
  managed_zone = "us-central1"
}

output "google_dns_record_sets" {
  value = data.google_dns_record_sets.pike
}
