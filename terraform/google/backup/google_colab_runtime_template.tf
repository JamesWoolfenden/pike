resource "google_colab_runtime_template" "my_template" {
  name         = "colab-runtime"
  display_name = "Runtime template basic"
  location     = "us-central1"

  machine_spec {
    machine_type = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}
