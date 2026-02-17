resource "google_compute_instance_template" "pike_autoscaler" {
  name         = "pike-autoscaler-template"
  machine_type = "e2-micro"

  disk {
    source_image = "debian-cloud/debian-11"
    auto_delete  = true
    boot         = true
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_group_manager" "pike_autoscaler" {
  name               = "pike-autoscaler-igm"
  base_instance_name = "pike-autoscaler"
  zone               = "us-central1-a"
  target_size        = 2

  version {
    instance_template = google_compute_instance_template.pike_autoscaler.id
  }
}

resource "google_compute_autoscaler" "pike" {
  name   = "pike-autoscaler"
  zone   = "us-central1-a"
  target = google_compute_instance_group_manager.pike_autoscaler.id

  autoscaling_policy {
    max_replicas    = 5
    min_replicas    = 2
    cooldown_period = 60

    cpu_utilization {
      target = 0.7
    }
  }
}
