
resource "google_pubsub_topic" "test_schema_topic" {
  name       = "test-schema-topic"
  depends_on = [google_pubsub_schema.test_schema]
  labels = {
    "name" : "test"
    "setup" : "terraform"
  }
  schema_settings {
    schema   = "projects/pike-gcp/schemas/${google_pubsub_schema.test_schema.name}"
    encoding = "JSON"
  }
}

resource "google_pubsub_subscription" "test_schema_sub" {
  name       = "test-schema-sub"
  topic      = google_pubsub_topic.test_schema_topic.name
  depends_on = [google_pubsub_topic.test_schema_topic]
}

#data google_project "pike" {
#
#}
