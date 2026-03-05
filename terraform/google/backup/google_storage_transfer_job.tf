resource "google_storage_transfer_job" "pike" {
  description = "Pike test transfer job"

  transfer_spec {
    gcs_data_source {
      bucket_name = google_storage_bucket.cache_bucket.name
    }

    gcs_data_sink {
      bucket_name = google_storage_bucket.cache_bucket.name
      path        = "destination/"
    }
  }

  schedule {
    schedule_start_date {
      year  = 2025
      month = 3
      day   = 5
    }
  }
}
