resource "aws_securityhub_configuration_policy" "example" {
  name        = "Example"
  description = "This is an example configuration policy"

  configuration_policy {
    service_enabled = true
    enabled_standard_arns = [
      "arn:aws:securityhub:us-east-1::standards/aws-foundational-security-best-practices/v/1.0.0",
      "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
    ]
    security_controls_configuration {
      disabled_control_identifiers = []
    }
  }

  depends_on = [aws_securityhub_organization_configuration.example]
}

resource "aws_securityhub_configuration_policy_association" "account_example" {
  target_id = "680235478471"
  policy_id = aws_securityhub_configuration_policy.example.id
}

resource "aws_securityhub_organization_configuration" "example" {
  auto_enable = true
}
