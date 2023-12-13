data "aws_caller_identity" "current" {}

resource "aws_opensearchserverless_access_policy" "example" {
  provider    = aws.central
  name        = "pike"
  type        = "data"
  description = "read and write permissions update"
  policy = jsonencode([
    {
      Rules = [
        {
          ResourceType = "index",
          Resource = [
            "index/example-collection/*"
          ],
          Permission = [
            "aoss:*"
          ]
        },
        {
          ResourceType = "collection",
          Resource = [
            "collection/example-collection"
          ],
          Permission = [
            "aoss:*"
          ]
        }
      ],
      Principal = [
        data.aws_caller_identity.current.arn
      ]
    }
  ])
}
