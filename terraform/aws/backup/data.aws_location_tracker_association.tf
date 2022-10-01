data "aws_location_tracker_association" "pike" {
  consumer_arn = "arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollectionConsumer"
  tracker_name = "example"
}
