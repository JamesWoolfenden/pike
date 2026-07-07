data "google_compute_target_https_proxy" "pike" {
  name = "pike"
}

output "google_compute_target_https_proxy" {
  value = data.google_compute_target_https_proxy.pike
}
