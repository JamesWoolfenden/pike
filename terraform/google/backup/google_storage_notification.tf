resource "google_storage_bucket" "pike_notification" {
  name     = "pike-notification-bucket-${random_id.bucket_suffix.hex}"
  location = "US"
}

resource "random_id" "bucket_suffix" {
  byte_length = 8
}

resource "google_pubsub_topic" "pike_notification" {
  name = "pike-storage-notification-topic"
}

resource "google_pubsub_topic_iam_member" "pike_notification" {
  topic  = google_pubsub_topic.pike_notification.name
  role   = "roles/pubsub.publisher"
  member = "serviceAccount:${data.google_storage_project_service_account.gcs_account.email_address}"
}

data "google_storage_project_service_account" "gcs_account" {
}

resource "google_storage_notification" "pike" {
  bucket         = google_storage_bucket.pike_notification.name
  payload_format = "JSON_API_V1"
  topic          = google_pubsub_topic.pike_notification.id
  event_types    = ["OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE"]

  depends_on = [google_pubsub_topic_iam_member.pike_notification]
}
