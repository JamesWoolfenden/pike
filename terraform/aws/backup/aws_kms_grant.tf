resource "aws_kms_grant" "pike" {
  name              = "my-grant"
  key_id            = data.aws_kms_key.pike.arn
  grantee_principal = data.aws_iam_role.pike.arn
  operations        = ["Encrypt", "Decrypt", "GenerateDataKey"]

  constraints {
    encryption_context_equals = {
      Department = "Finance"
    }
  }
}

data "aws_iam_role" "pike" {
  name = "pike-test-role"
}

data "aws_kms_key" "pike" {
  key_id = "34cdce9a-2322-427c-91bb-b572f435c032"
}
