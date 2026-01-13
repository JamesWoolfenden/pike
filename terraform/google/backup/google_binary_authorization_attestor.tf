resource "google_binary_authorization_attestor" "attestor" {
  name = "test-attestor"
  attestation_authority_note {
    note_reference = google_container_analysis_note.note.name
  }
}


resource "google_container_analysis_note" "note" {
  name = "attestor-note"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}
