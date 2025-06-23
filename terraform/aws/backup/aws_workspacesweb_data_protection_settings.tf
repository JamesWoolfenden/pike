
resource "aws_workspacesweb_data_protection_settings" "example" {
  display_name         = "example-complete"
  description          = "Complete example data protection settings"
  customer_managed_key = aws_kms_key.example.arn

  additional_encryption_context = {
    Environment = "Production"
  }

  inline_redaction_configuration {
    global_confidence_level = 2
    global_enforced_urls    = ["https://example.com", "https://test.example.com"]
    global_exempt_urls      = ["https://exempt.example.com"]

    inline_redaction_pattern {
      built_in_pattern_id = "ssn"
      confidence_level    = 3
      enforced_urls       = ["https://pattern1.example.com"]
      exempt_urls         = ["https://exempt-pattern1.example.com"]
      redaction_place_holder {
        redaction_place_holder_type = "CustomText"
        redaction_place_holder_text = "REDACTED-SSN"
      }
    }

    inline_redaction_pattern {
      custom_pattern {
        pattern_name        = "CustomPattern"
        pattern_regex       = "/\\d{3}-\\d{2}-\\d{4}/g"
        keyword_regex       = "/SSN|Social Security/gi"
        pattern_description = "Custom SSN pattern"
      }
      redaction_place_holder {
        redaction_place_holder_type = "CustomText"
        redaction_place_holder_text = "REDACTED-CUSTOM"
      }
    }
  }

  tags = {
    Name = "example-data-protection-settings"
  }
}
