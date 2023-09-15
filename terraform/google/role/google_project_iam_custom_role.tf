resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    // google_pubsub_lite_topic
    "cloudkms.cryptoKeys.getIamPolicy",
    "cloudkms.cryptoKeys.setIamPolicy"
  ]
}
