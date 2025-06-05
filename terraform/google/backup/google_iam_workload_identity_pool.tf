resource "google_iam_workload_identity_pool" "pike" {
  workload_identity_pool_id = "example-pike-pool2"
  display_name              = "Name of pool"
  description               = "Identity pool for automated test 2"
  disabled                  = true
}
