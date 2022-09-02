resource "aws_glue_data_catalog_encryption_settings" "example" {
  data_catalog_encryption_settings {
    connection_password_encryption {
      aws_kms_key_id                       = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
      return_connection_password_encrypted = true
    }

    encryption_at_rest {
      catalog_encryption_mode = "SSE-KMS"
      sse_aws_kms_key_id      = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
    }
  }
}
