resource "aws_location_geofence_collection" "pike" {
  collection_name = "pikey"
  description     = "pike"
  kms_key_id      = "arn:aws:kms:eu-west-2:680235478471:key/554dbedc-3cf9-4aec-b621-9c4387bed449"
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
