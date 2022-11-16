resource "aws_cloudfront_field_level_encryption_profile" "pike" {
  comment = "test comment pike"
  name    = "test_profile"

  encryption_entities {
    items {
      public_key_id = aws_cloudfront_public_key.pike.id
      provider_id   = "test_provider"

      field_patterns {
        items = ["DateOfBirth"]
      }
    }
  }
}
