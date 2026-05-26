resource "google_pubsub_lite_topic" "pike" {
  name = "pike"
  partition_config {
    count = 1
    capacity {
      publish_mib_per_sec   = 4
      subscribe_mib_per_sec = 8
    }
  }

  retention_config {
    per_partition_bytes = 32212254720
  }

  reservation_config {
    throughput_reservation = google_pubsub_lite_reservation.pike.name
  }
}
