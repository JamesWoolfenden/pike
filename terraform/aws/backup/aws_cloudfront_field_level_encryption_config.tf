resource "aws_cloudfront_field_level_encryption_config" "pike" {
  comment = "test comment pike"


  content_type_profile_config {

    forward_when_content_type_is_unknown = true

    content_type_profiles {
      items {
        content_type = "application/x-www-form-urlencoded"
        format       = "URLEncoded"
        profile_id   = aws_cloudfront_field_level_encryption_profile.pike.id
      }
    }
  }

  query_arg_profile_config {
    forward_when_query_arg_profile_is_unknown = true

    query_arg_profiles {
      items {
        profile_id = aws_cloudfront_field_level_encryption_profile.pike.id
        query_arg  = "Arg1"
      }
    }
  }
}
