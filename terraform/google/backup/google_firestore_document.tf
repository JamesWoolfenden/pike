resource "google_firestore_document" "pike" {
  database    = google_firestore_database.pike.name
  collection  = "somenewcollection"
  document_id = "my-doc-id"
  fields      = "{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}"
}
