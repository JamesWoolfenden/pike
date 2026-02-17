resource "google_compute_instance_group_membership" "pike" {
  instance       = google_compute_instance.pike.self_link
  instance_group = google_compute_instance_group.pike.name
  zone           = "us-central1-a"
}
