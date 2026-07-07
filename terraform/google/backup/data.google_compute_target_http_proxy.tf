data "google_compute_target_http_proxy" "pike" {
  name = "pike"
}

output "google_compute_target_http_proxy" {
  value = data.google_compute_target_http_proxy.pike
}
