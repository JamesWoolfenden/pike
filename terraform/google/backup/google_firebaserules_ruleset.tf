resource "google_firebaserules_ruleset" "firestore" {
  source {
    files {
      content = "service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }"
      name    = "firestore.rules"
    }
  }

  project = google_firebase_project.pike.project
  depends_on = [
    google_firebase_project.pike
  ]
}
