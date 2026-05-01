data "google_netblock_ip_ranges" "pike" {}

output "google_netblock_ip_ranges" {
  value = data.google_netblock_ip_ranges.pike
}
