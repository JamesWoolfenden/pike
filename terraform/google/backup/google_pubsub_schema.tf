resource "google_pubsub_schema" "test_schema" {
  name       = "test-schema"
  type       = "PROTOCOL_BUFFER"
  definition = "syntax = 'proto3'; message TestSchema { string one = 1; string two = 2; string three = 3; }"
}
