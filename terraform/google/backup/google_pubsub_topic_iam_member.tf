resource "google_pubsub_topic_iam_member" "pike" {
  topic  = "projects/pike-gcp/topics/pike"
  member = "user:anonymous@gmail.com"
  role   = "roles/viewer"
}
