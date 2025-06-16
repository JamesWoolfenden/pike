data "google_pubsub_schema_iam_policy" "pike" {
  schema = google_pubsub_schema.example.id
}

output "google_pubsub_schema_iam_policy" {
  value = data.google_pubsub_schema_iam_policy.pike
}

resource "google_pubsub_schema" "example" {
  name       = "example-schema"
  type       = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}
