data "google_compute_instance_template_iam_policy" "pike" {
  name = "pike"
}

output "google_compute_instance_template_iam_policy" {
  value = data.google_compute_instance_template_iam_policy.pike
}
