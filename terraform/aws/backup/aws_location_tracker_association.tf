resource "aws_location_tracker_association" "pike" {
  consumer_arn = aws_location_geofence_collection.pike.collection_arn
  tracker_name = aws_location_tracker.pike.tracker_name
}
