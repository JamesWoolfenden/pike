resource "google_project_iam_member" "owner" {
  project = "p"
  role    = "roles/owner"
  member  = "user:a@example.com"
}

resource "google_project_iam_binding" "editor_public" {
  project = "p"
  role    = "roles/editor"
  members = ["allUsers", "user:a@example.com"]
}

resource "google_folder_iam_member" "impersonate" {
  folder = "folders/123"
  role   = "roles/iam.serviceAccountTokenCreator"
  member = "serviceAccount:x@p.iam.gserviceaccount.com"
}

resource "google_project_iam_policy" "authoritative" {
  project     = "p"
  policy_data = data.google_iam_policy.x.policy_data
}

resource "google_project_iam_custom_role" "esc" {
  project     = "p"
  role_id     = "esc"
  title       = "esc"
  permissions = ["compute.instances.get", "iam.roles.create"]
}

resource "google_service_account_iam_member" "ok_scope" {
  service_account_id = "projects/p/serviceAccounts/x"
  role               = "roles/iam.serviceAccountTokenCreator"
  member             = "user:a@example.com"
}
