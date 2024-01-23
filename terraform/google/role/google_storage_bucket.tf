
resource "google_storage_bucket" "default" {
  name          = "terraform-pike-bucket-tfstate"
  force_destroy = false
  location      = "EUROPE-WEST2"
  storage_class = "STANDARD"
  versioning {
    enabled = true
  }
  encryption {
    default_kms_key_name = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/state"
  }

  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      age                        = 0
      days_since_custom_time     = 0
      days_since_noncurrent_time = 0
      matches_prefix             = []
      matches_storage_class      = []
      matches_suffix             = []
      num_newer_versions         = 1
      with_state                 = "ARCHIVED"
    }
  }
  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      age                        = 0
      days_since_custom_time     = 0
      days_since_noncurrent_time = 1
      matches_prefix             = []
      matches_storage_class      = []
      matches_suffix             = []
      num_newer_versions         = 0
      with_state                 = "ANY"
    }
  }

  timeouts {}

}
