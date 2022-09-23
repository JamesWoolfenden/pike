resource "aws_ecr_repository_policy" "pike" {
  policy     = data.aws_iam_policy_document.allowlocals.json
  repository = "pike"
}

data "aws_iam_policy_document" "allowlocals" {
  statement {
    sid = "AllowPushPull"

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:DescribeRepositories",
      "ecr:GetRepositoryPolicy",
      "ecr:ListImages",
      "ecr:DeleteRepository",
      "ecr:BatchDeleteImage",
      "ecr:SetRepositoryPolicy",
      "ecr:DeleteRepositoryPolicy"
    ]

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.ecr.account_id}:root"]
    }
  }
}

data "aws_caller_identity" "ecr" {}
