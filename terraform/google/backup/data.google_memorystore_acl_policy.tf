data "google_memorystore_acl_policy" "pike" {
  acl_policy_id = "pike-test-acl1"
}

output "google_memorystore_acl_policy" {
  value = data.google_memorystore_acl_policy.pike
}