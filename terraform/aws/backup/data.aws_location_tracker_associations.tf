data "aws_location_tracker_associations" "pike" {
  tracker_name = "pike"
}

output "assoc" {
  value = data.aws_location_tracker_associations.pike
}
