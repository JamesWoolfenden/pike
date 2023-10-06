data "google_secret_manager_secret_iam_policy" "pike" {
  project   = data.google_secret_manager_secret.pike.project
  secret_id = data.google_secret_manager_secret.pike.secret_id
}
