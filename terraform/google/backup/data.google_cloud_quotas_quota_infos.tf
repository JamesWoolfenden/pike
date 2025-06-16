data "google_cloud_quotas_quota_infos" "pike" {
  service = "pike"
  parent  = "pike"
}

output "google_cloud_quotas_quota_infos" {
  value = data.google_cloud_quotas_quota_infos.pike
}
