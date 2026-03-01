# Region Commitment - committed use discounts for compute resources
resource "google_compute_region_commitment" "pike" {
  name    = "pike-region-commitment"
  region  = "us-central1"
  project = "pike-477416"

  plan = "TWELVE_MONTH"

  resources {
    type   = "VCPU"
    amount = "4"
  }

  resources {
    type   = "MEMORY"
    amount = "16384"
  }
}
