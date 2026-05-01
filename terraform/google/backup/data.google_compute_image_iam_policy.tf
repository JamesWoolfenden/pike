data "google_compute_image_iam_policy" "pike" {
  image = "ubuntu"
}

output "google_compute_image_iam_policy" {
  value = data.google_compute_image_iam_policy.pike
}
