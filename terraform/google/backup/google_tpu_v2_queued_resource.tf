resource "google_tpu_v2_queued_resource" "qr" {
  provider = google-beta

  name = "test-qr"
  zone = "us-central1-c"

  tpu {
    node_spec {
      parent  = "projects/pike-412922/locations/us-central1-c"
      node_id = "test-tpu"
      node {
        runtime_version  = "tpu-vm-tf-2.13.0"
        accelerator_type = "v2-8"
        description      = "Text description of the TPU."
      }
    }
  }
}
