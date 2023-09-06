data "aws_caller_identity" "current" {}

data "aws_iam_principal_policy_simulation" "s3_object_access" {
  action_names = [
    "s3:GetObject",
    "s3:PutObject",
    "s3:DeleteObject",
  ]
  policy_source_arn = data.aws_caller_identity.current.arn
  resource_arns     = ["arn:aws:s3:::my-test-bucket"]

  # The "lifecycle" and "postcondition" block types are part of
  # the main Terraform language, not part of this data source.
  lifecycle {
    postcondition {
      condition     = self.all_allowed
      error_message = <<EOT
        Given AWS credentials do not have sufficient access to manage ${join(", ", self.resource_arns)}.
      EOT
    }
  }
}
