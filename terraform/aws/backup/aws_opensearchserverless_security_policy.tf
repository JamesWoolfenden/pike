resource "aws_opensearchserverless_security_policy" "pike" {
  provider    = aws.central
  name        = "pike"
  type        = "encryption"
  description = "encryption security policy for example-collection update"
  policy = jsonencode({
    Rules = [
      {
        Resource = [
          "collection/pike"
        ],
        ResourceType = "collection"
      }
    ],
    AWSOwnedKey = true
  })

}
