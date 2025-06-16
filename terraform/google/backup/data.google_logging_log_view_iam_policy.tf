data "google_logging_log_view_iam_policy" "pike" {
  bucket = "anyoldshet"
  parent = "pike"
  name   = "pike"
}

output "google_logging_log_view_iam_policy" {
  value = data.google_logging_log_view_iam_policy.pike
}
