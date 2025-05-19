resource "aws_dataexchange_revision_assets" "pike" {
  data_set_id = "exampleidmustbelongerthan30characters"

  asset {
    create_s3_data_access_from_s3_bucket {
      asset_source {
        bucket = "example-bucket"
      }
    }
  }

  tags = {
    pike        = "permission"
    Environment = "Production"
  }
}
