resource "aws_glue_schema" "pike" {
  schema_name       = "pike"
  registry_arn      = "arn:aws:glue:eu-west-2:680235478471:registry/pike-registry"
  data_format       = "AVRO"
  compatibility     = "NONE"
  schema_definition = "{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}"
  description       = "pike schema mod"
  tags = {
    pike   = "permissions"
    delete = "me"
    tag    = "again"
  }
}
