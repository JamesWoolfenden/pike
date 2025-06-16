data "google_colab_runtime_template_iam_policy" "pike" {
  runtime_template = "pike"
}

output "google_colab_runtime_template_iam_policy" {
  value = data.google_colab_runtime_template_iam_policy.pike
}
