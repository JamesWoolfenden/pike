
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "logging.logEntries.list" # placeholder — provider rejects an empty list (min 1 item); unrelated to any type under test
  ]
}
