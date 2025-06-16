data "google_folder_iam_policy" "pike" {
  folder = "folders/pike"
}

output "google_folder_iam_policy" {
  value = data.google_folder_iam_policy.pike
}
