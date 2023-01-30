resource "aws_s3_bucket_inventory" "pike" {
  bucket                   = "config-store-jgw"
  included_object_versions = "Current"
  name                     = "EntireBucketDaily"
  enabled = true
  destination {
    bucket {
      account_id = "680235478471"
      bucket_arn = "arn:aws:s3:::test-inventory680235478471"
      format     = "ORC"
      encryption {
        sse_s3 {}
      }
    }
  }
  schedule {
    frequency = "Daily"
  }
}

#data aws_s3_bucket "config" {
#
#  bucket = "config-store-jgw"
#}