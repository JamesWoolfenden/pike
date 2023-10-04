resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_kms_secret
    "cloudkms.cryptoKeyVersions.useToDecrypt",
    //google_kms_key_ring_iam_policy
    "cloudkms.keyRings.getIamPolicy",
    //google_pubsub_subscription
    "pubsub.subscriptions.get",
    //topics
    "pubsub.topics.get",
    "pubsub.topics.getIamPolicy",
    "pubsub.subscriptions.getIamPolicy",

    //google_kms_secret_asymmetric
    "cloudkms.cryptoKeyVersions.useToDecrypt"

  ]
}
