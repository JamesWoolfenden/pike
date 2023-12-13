resource "aws_ecrpublic_repository" "foo" {
  provider        = aws.central
  repository_name = "bar"
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}

data "aws_iam_policy_document" "foopolicy" {
  statement {
    sid    = "new policy"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["680235478471"]
    }

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
      "ecr:DeleteRepositoryPolicy",
    ]
  }
}

resource "aws_ecrpublic_repository_policy" "foopolicy" {
  provider        = aws.central
  repository_name = aws_ecrpublic_repository.foo.repository_name
  policy          = data.aws_iam_policy_document.foopolicy.json
}
