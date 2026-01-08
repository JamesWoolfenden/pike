resource "google_iam_folders_policy_binding" "pike" {
  folder            = google_folder.folder.folder_id
  location          = "global"
  display_name      = "binding for all principals in the folder"
  policy_kind       = "PRINCIPAL_ACCESS_BOUNDARY"
  policy_binding_id = "binding-for-all-folder-principals"
  policy            = "{}"
  target {
    principal_set = "//cloudresourcemanager.googleapis.com/folders/${google_folder.folder.folder_id}"
  }
  depends_on = [time_sleep.wait_120s]
}

resource "google_folder" "folder" {
  display_name        = "my folder"
  parent              = "organizations/123456789"
  deletion_protection = false
}

resource "time_sleep" "wait_120s" {
  depends_on      = [google_folder.folder]
  create_duration = "120s"
}
