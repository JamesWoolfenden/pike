data "google_compute_forwarding_rules" "pike" {
}

output "google_compute_forwarding_rules" {
  value = data.google_compute_forwarding_rules.pike
}
