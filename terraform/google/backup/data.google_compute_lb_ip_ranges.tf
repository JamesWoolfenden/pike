data "google_compute_lb_ip_ranges" "pike" {}


output "google_compute_lb_ip_ranges" {
  value = data.google_compute_lb_ip_ranges.pike
}
