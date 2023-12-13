data "google_netblock_ip_ranges" "pike" {}

output "ranges" {
  value = data.google_netblock_ip_ranges.pike
}
