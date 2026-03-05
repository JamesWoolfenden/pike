data "google_firestore_document" "pike" {
  document_id = "pike"
  collection  = "pike"
  database    = "pike"
}

output "google_firestore_document" {
  value = data.google_firestore_document.pike
}
