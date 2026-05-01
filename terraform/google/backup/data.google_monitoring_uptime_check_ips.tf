data "google_monitoring_uptime_check_ips" "pike" {}

output "google_monitoring_uptime_check_ips" {
  value = data.google_monitoring_uptime_check_ips.pike
}
