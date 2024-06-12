resource "aws_sagemaker_model_package_group_policy" "pike" {
  model_package_group_name = aws_sagemaker_model_package_group.pike.model_package_group_name
  resource_policy          = jsonencode(jsondecode(data.aws_iam_policy_document.example.json))
}

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "example" {
  statement {
    sid       = "AddPermModelPackageGroup"
    actions   = ["sagemaker:DescribeModelPackage", "sagemaker:ListModelPackages"]
    resources = [aws_sagemaker_model_package_group.pike.arn]
    principals {
      identifiers = [data.aws_caller_identity.current.account_id]
      type        = "AWS"
    }
  }
}
