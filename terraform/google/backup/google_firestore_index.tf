resource "google_firestore_index" "pike" {

  database   = google_firestore_database.pike.name
  collection = "atestcollection"

  fields {
    field_path = "name"
    order      = "ASCENDING"
  }

  fields {
    field_path = "description"
    order      = "DESCENDING"
  }
}
