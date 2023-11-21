resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_storage_bucket_access_control, google_storage_default_object_access_control,google_storage_default_object_acl
    "storage.buckets.get",
    "storage.buckets.update",
    //google_storage_object_access_control
    "storage.objects.update",
    //google_storage_bucket_iam_member
    "storage.buckets.getIamPolicy",
    //google_storage_bucket_iam_policy
    "storage.buckets.setIamPolicy",
    //google_storage_object_access_control
    "storage.objects.get",
    "storage.objects.getIamPolicy",
    "storage.objects.setIamPolicy",
    //google_storage_insights_report_config
    "storageinsights.reportConfigs.create",
    //google_storage_hmac_key
    "storage.hmacKeys.create",
    "storage.hmacKeys.get",
    "storage.hmacKeys.update",
    "storage.hmacKeys.delete"
  ]
}
