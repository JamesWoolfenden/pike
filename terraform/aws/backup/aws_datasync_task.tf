resource "aws_datasync_task" "pike" {
  destination_location_arn = aws_datasync_location_s3.destination_pike.arn
  name                     = "example"
  source_location_arn      = aws_datasync_location_nfs.source_pike.arn

  options {
    bytes_per_second = -1
  }

  tags = {
    pike = "permissions"
  }
}
