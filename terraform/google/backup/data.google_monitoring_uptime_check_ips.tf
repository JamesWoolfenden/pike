data "google_monitoring_uptime_check_ips" "pike" {}

output "ips" {
  value = data.google_monitoring_uptime_check_ips.pike
}
