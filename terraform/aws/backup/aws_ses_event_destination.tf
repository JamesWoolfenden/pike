resource "aws_ses_event_destination" "pike" {
  name                   = "event-destination-cloudwatch"
  configuration_set_name = aws_ses_configuration_set.pike.name
  enabled                = false
  matching_types         = ["bounce", "send"]

  cloudwatch_destination {
    default_value  = "default"
    dimension_name = "dimension"
    value_source   = "emailHeader"
  }
}
