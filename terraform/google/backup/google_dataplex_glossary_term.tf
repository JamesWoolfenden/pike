resource "google_dataplex_glossary_term" "pike" {
  parent      = "projects/${google_dataplex_glossary.term_test_id.project}/locations/us-central1/glossaries/${google_dataplex_glossary.term_test_id.glossary_id}"
  glossary_id = google_dataplex_glossary.term_test_id.glossary_id
  location    = "us-central1"
  term_id     = "tf-test-term-basic"
}
