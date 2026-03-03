resource "google_logging_saved_query" "pike" {
  name         = "pike-saved-query"
  parent       = "projects/pike-477416"
  location     = "global"
  visibility   = "PRIVATE"
  display_name = "Pike Test Query"
  description  = "Pike test saved query - updated"

  logging_query {
    filter = "severity >= ERROR"
  }
}
