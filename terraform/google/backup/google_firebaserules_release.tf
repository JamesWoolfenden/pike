resource "google_firebaserules_release" "primary" {
  name         = "cloud.firestore"
  ruleset_name = "projects/pike-gcp/rulesets/${google_firebaserules_ruleset.firestore.name}"
  project      = google_firebase_project.pike.project

  lifecycle {
    replace_triggered_by = [
      google_firebaserules_ruleset.firestore
    ]
  }
  depends_on = [
    google_firebase_project.pike
  ]
}
