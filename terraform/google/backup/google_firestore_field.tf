resource "google_firestore_field" "pike" {

  database   = google_firestore_database.pike.name
  collection = "chatrooms_"
  field      = "basic"

  index_config {
    indexes {
      order       = "ASCENDING"
      query_scope = "COLLECTION_GROUP"
    }
    indexes {
      array_config = "CONTAINS"
    }
  }
}
