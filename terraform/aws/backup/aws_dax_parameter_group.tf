resource "aws_dax_parameter_group" "pike" {
  name = "pike"
  parameters {
    name  = "query-ttl-millis"
    value = "100000"
  }

  parameters {
    name  = "record-ttl-millis"
    value = "100000"
  }
}
