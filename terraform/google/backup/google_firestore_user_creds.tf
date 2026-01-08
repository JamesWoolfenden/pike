resource "google_firestore_user_creds" "pike" {
  database = google_firestore_database.pike.name
  name     = "my-username"
}
