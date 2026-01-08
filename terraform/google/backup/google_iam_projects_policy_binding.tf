data "google_project" "project" {
  provider = google
}



resource "time_sleep" "wait_60_seconds" {
  create_duration = "60s"
}

resource "google_iam_projects_policy_binding" "binding-for-all-project-principals" {
  depends_on        = [time_sleep.wait_60_seconds]
  project           = data.google_project.project.project_id
  location          = "global"
  display_name      = "binding for all principals in the project"
  policy_kind       = "PRINCIPAL_ACCESS_BOUNDARY"
  policy_binding_id = "binding-for-all-project-principals"
  policy            = "{}"
  target {
    principal_set = "//cloudresourcemanager.googleapis.com/projects/${data.google_project.project.project_id}"
  }
}
