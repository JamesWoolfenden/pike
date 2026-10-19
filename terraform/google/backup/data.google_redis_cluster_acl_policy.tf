data "google_redis_cluster_acl_policy" "pike" {
  acl_policy_id = "pike-test-acl1"
}

output "google_redis_cluster_acl_policy" {
  value = data.google_redis_cluster_acl_policy.pike
}