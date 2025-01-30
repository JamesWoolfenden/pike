resource "aws_securityhub_standards_subscription" "pike" {
  standards_arn = "arn:aws:securityhub:us-east-1:123456789012:subscription/cis-aws-foundations-benchmark/v/1.2.0"
}
