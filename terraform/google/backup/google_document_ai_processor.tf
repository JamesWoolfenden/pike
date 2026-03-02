# Document AI Processor
resource "google_document_ai_processor" "pike" {
  display_name = "pike-test-processor"
  type         = "OCR_PROCESSOR"
  location     = "us"
  project      = "pike-477416"
}
