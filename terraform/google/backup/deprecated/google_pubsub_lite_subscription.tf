resource "google_pubsub_lite_subscription" "pike" {
  name  = "pike"
  topic = google_pubsub_lite_topic.pike.name
  delivery_config {
    delivery_requirement = "DELIVER_AFTER_STORED"
  }
}
