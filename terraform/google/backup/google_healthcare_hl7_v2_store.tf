# Healthcare HL7 V2 Store
resource "google_healthcare_dataset" "pike" {
  name     = "pike-dataset"
  location = "us-central1"
  project  = "pike-477416"
}

resource "google_healthcare_hl7_v2_store" "pike" {
  name    = "pike-hl7-store"
  dataset = google_healthcare_dataset.pike.id

  labels = {
    environment = "test"
  }
}
