data "google_compute_lb_ip_ranges" "pike" {}


output "ranges" {
  value = data.google_compute_lb_ip_ranges.pike
}
