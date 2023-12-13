resource "aws_opensearchserverless_collection" "pike" {
  provider = aws.central
  name     = "pike"
  tags = {
    pike    = "permissions"
    another = "tag"
  }

  depends_on = [
  aws_opensearchserverless_security_policy.pike]
}
