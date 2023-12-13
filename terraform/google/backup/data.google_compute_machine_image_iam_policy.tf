data "google_compute_machine_image_iam_policy" "pike" {
  provider      = google-beta
  machine_image = "pike"
}
