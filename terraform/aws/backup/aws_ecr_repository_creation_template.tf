resource "aws_ecr_repository_creation_template" "pike" {
  prefix               = "example"
  description          = "An example template"
  image_tag_mutability = "IMMUTABLE"
  custom_role_arn      = "arn:aws:iam::123456789012:role/example"

  applied_for = [
    "PULL_THROUGH_CACHE",
  ]

  encryption_configuration {
    encryption_type = "AES256"
  }

  repository_policy = data.aws_iam_policy_document.pike.json

  lifecycle_policy = <<EOT
{
  "rules": [
    {
      "rulePriority": 1,
      "description": "Expire images older than 14 days",
      "selection": {
        "tagStatus": "untagged",
        "countType": "sinceImagePushed",
        "countUnit": "days",
        "countNumber": 14
      },
      "action": {
        "type": "expire"
      }
    }
  ]
}
EOT

  resource_tags = {
    Foo = "Bar"
  }
}
