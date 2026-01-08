resource "google_iam_workforce_pool_provider_scim_tenant" "tenant" {
  location          = "global"
  workforce_pool_id = google_iam_workforce_pool.pool.workforce_pool_id
  provider_id       = google_iam_workforce_pool_provider.provider.provider_id
  scim_tenant_id    = "example-tenant"
  display_name      = "SCIM Tenant display Name"
  description       = "A SCIM Tenant for IAM Workforce Pool Provider"
  claim_mapping = {
    "google.subject" = "user.externalId",
    "google.group"   = "group.externalId"
  }
}
