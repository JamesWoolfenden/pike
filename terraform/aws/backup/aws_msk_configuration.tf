resource "aws_msk_configuration" "pike" {
  kafka_versions = ["3.2.0"]
  name           = "pike"

  server_properties = <<PROPERTIES
auto.create.topics.enable = true
delete.topic.enable = true
PROPERTIES
}
