data "google_compute_image_iam_policy" "pike" {
  image = "ubuntu"
}

output "policy" {
  value = data.google_compute_image_iam_policy.pike
}
