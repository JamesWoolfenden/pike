resource "aws_s3_bucket_intelligent_tiering_configuration" "pike" {
  bucket = "pike-680235478471"
  name   = "ImportantBlueDocuments"

  status = "Disabled"

  filter {
    prefix = "documents/"

    tags = {
      priority = "high"
      class    = "blue"
    }
  }

  tiering {
    access_tier = "ARCHIVE_ACCESS"
    days        = 125
  }
}
