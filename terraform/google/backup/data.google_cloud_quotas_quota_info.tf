data "google_cloud_quotas_quota_info" "pike" {
  parent   = "pike"
  quota_id = "pike"
  service  = "pike"
}

output "google_cloud_quotas_quota_info" {
  value = data.google_cloud_quotas_quota_info.pike
}
