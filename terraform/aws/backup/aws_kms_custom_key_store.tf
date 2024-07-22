resource "aws_kms_custom_key_store" "pike" {
  custom_key_store_name    = "pike"
  cloud_hsm_cluster_id     = ""
  key_store_password       = ""
  trust_anchor_certificate = ""
}
